package service

import (
	"context"
	"geemod/jobs-api/dto"
)

type CompanyService interface {
	Create(ctx context.Context, request dto.CompanyCreateRequest) dto.CompanyResponse
	Update(ctx context.Context, request dto.CompanyUpdateRequest) dto.CompanyResponse
	FindByID(ctx context.Context, categoryID string) dto.CompanyResponse
}
