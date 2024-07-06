package service

import (
	"context"
	"database/sql"
	"geemod/jobs-api/dto"
	"geemod/jobs-api/exception"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/model"
	"geemod/jobs-api/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type JobsServiceImpl struct{
	JobsRepository repository.JobsRepository
	CompanyRepository repository.CompanyRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewJobsService(jobsRepository repository.JobsRepository, companyRepository repository.CompanyRepository, DB *sql.DB, validate *validator.Validate)JobsService{
	return &JobsServiceImpl{
		JobsRepository: jobsRepository,
		CompanyRepository: companyRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *JobsServiceImpl) Create(ctx context.Context, request dto.JobsCreateRequest) dto.JobsResponse{
	
	err := service.Validate.Struct(request)
	helper.SendPanicError(err)

	tx, err := service.DB.Begin()
	helper.SendPanicError(err)
	defer helper.CommitOrRollback(tx)

	company , err  := service.CompanyRepository.FindById(ctx, tx, request.CompanyId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	newUUID := uuid.New()
	uuidString := newUUID.String();
	
	jobs := model.Jobs{
		Id: uuidString,
		CompanyId: company.Id,
		Company: &company,
		Tittle: request.Tittle,
		Description: request.Description,
		CreateAt: time.Now(),
	}

	jobs = service.JobsRepository.Save(ctx, tx, jobs)

	return helper.ToJobsRespose(jobs)
	
} 

func (service *JobsServiceImpl) FindAll(ctx context.Context) []dto.JobsResponse{
	
	tx, err  := service.DB.Begin()
	helper.SendPanicError(err)
	defer helper.CommitOrRollback(tx)

	jobs := service.JobsRepository.FindAll(ctx, tx)

	return helper.ToJobsListResponse(jobs)

}