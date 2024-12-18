package repository

import (
	"context"
	"database/sql"
	"geemod/jobs-api/model"
)

type CompanyRepository interface {
	Save(ctx context.Context, tx *sql.Tx, company model.Company) model.Company
	Update(ctx context.Context, tx *sql.Tx, company model.Company) model.Company
	FindById(ctx context.Context, tx *sql.Tx, companyId string) (model.Company, error)
}
