package service

import (
	"context"
	"database/sql"
	"geemod/jobs-api/dto"
	"geemod/jobs-api/exception"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/model"
	"geemod/jobs-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CompanyServiceImpl struct {
	CompanyRepository repository.CompanyRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCompanyService(companyRepository repository.CompanyRepository, DB *sql.DB, validate *validator.Validate) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: companyRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CompanyServiceImpl) Create(ctx context.Context, request dto.CompanyCreateRequest) dto.CompanyResponse {

	err := service.Validate.Struct(request)
	helper.SendPanicError(err)

	tx, err := service.DB.Begin()
	helper.SendPanicError(err)
	defer helper.CommitOrRollback(tx)

	newUUID := uuid.New()
	uuidString := newUUID.String()

	company := model.Company{
		Id:   uuidString,
		Name: request.Name,
	}

	company = service.CompanyRepository.Save(ctx, tx, company)

	return helper.ToCompanyResponse(company)

}

func (service *CompanyServiceImpl) Update(ctx context.Context, request dto.CompanyUpdateRequest) dto.CompanyResponse {

	err := service.Validate.Struct(request)
	helper.SendPanicError(err)

	tx, err := service.DB.Begin()
	helper.SendPanicError(err)
	defer helper.CommitOrRollback(tx)

	company, err := service.CompanyRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	company.Name = request.Name
	company = service.CompanyRepository.Update(ctx, tx, company)

	return helper.ToCompanyResponse(company)

}

func (service *CompanyServiceImpl) FindByID(ctx context.Context, categoryID string) dto.CompanyResponse {
	panic("impl mee")
}
