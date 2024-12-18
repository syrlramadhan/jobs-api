package helper

import (
	"geemod/jobs-api/dto"
	"geemod/jobs-api/model"
)

func ToCompanyResponse(company model.Company) dto.CompanyResponse {
	return dto.CompanyResponse{
		Id:   company.Id,
		Name: company.Name,
	}
}

func ToJobsRespose(jobs model.Jobs) dto.JobsResponse {
	return dto.JobsResponse{
		Id:          jobs.Id,
		CompanyId:   jobs.CompanyId,
		CompanyName: jobs.Company.Name,
		Tittle:      jobs.Tittle,
		Description: jobs.Description,
		CreatedAt:   jobs.CreateAt,
	}
}

func ToJobsListResponse(jobs []model.Jobs) []dto.JobsResponse {
	var jobsResponses []dto.JobsResponse
	for _, job := range jobs {
		jobsResponses = append(jobsResponses, ToJobsRespose(job))
	}
	return jobsResponses
}
