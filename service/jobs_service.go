package service

import (
	"context"
	"geemod/jobs-api/dto"
)

type JobsService interface{
	Create(ctx context.Context, request dto.JobsCreateRequest) dto.JobsResponse
	FindAll(ctx context.Context, companyName string, limit int) []dto.JobsResponse
	Count(ctx context.Context, companyName string) int
}