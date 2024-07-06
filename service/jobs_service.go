package service

import (
	"context"
	"geemod/jobs-api/dto"
)

type JobsService interface{
	Create(ctx context.Context, request dto.JobsCreateRequest) dto.JobsResponse
	FindAll(ctx context.Context) []dto.JobsResponse
}