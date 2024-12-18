package repository

import (
	"context"
	"database/sql"
	"errors"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/model"
)

type CompanyRepositoryImpl struct {
}

func NewCompanyRepository() CompanyRepository {
	return CompanyRepositoryImpl{}
}

func (c CompanyRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, company model.Company) model.Company {

	SQL := "insert into mst_companies(id, name) values(?, ?)"
	_, err := tx.ExecContext(ctx, SQL, company.Id, company.Name)
	helper.SendPanicError(err)

	return company
}

func (c CompanyRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, company model.Company) model.Company {

	SQL := "update mst_companies set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, company.Name, company.Id)
	helper.SendPanicError(err)

	return company
}

func (c CompanyRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, companyId string) (model.Company, error) {

	SQL := "select id, name from mst_companies where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, companyId)
	helper.SendPanicError(err)

	defer rows.Close()
	company := model.Company{}
	if rows.Next() {
		err := rows.Scan(&company.Id, &company.Name)
		helper.SendPanicError(err)
		return company, nil
	} else {
		return company, errors.New("company is not found")
	}

}
