package repository

import (
	"context"
	"database/sql"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/model"
)

type JobsRepositoryImpl struct {
}

func NewJobsRepository() JobsRepository {
	return JobsRepositoryImpl{}
}

func (c JobsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, jobs model.Jobs) model.Jobs {
	SQL := "insert into mst_jobs(id, company_id, tittle, description) values(?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, jobs.Id, jobs.CompanyId, jobs.Tittle, jobs.Description)
	helper.SendPanicError(err)

	return jobs
}

func (c JobsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, companyName string, limit int) []model.Jobs {

	SQL := "select j.id, j.tittle, j.description, c.name from mst_jobs j JOIN mst_companies c on j.company_id = c.id where c.name like ? limit ?"
	rows, err := tx.Query(SQL, "%"+companyName+"%", limit)
	helper.SendPanicError(err)
	defer rows.Close()

	var jobs []model.Jobs
	for rows.Next() {
		job := model.Jobs{}
		job.Company = &model.Company{}
		err := rows.Scan(&job.Id, &job.Tittle, &job.Description, &job.Company.Name)
		helper.SendPanicError(err)
		jobs = append(jobs, job)
	}

	return jobs
}

func (c JobsRepositoryImpl) Count(ctx context.Context, tx *sql.Tx, companyName string) int {

	var result int

	SQL := "select count(*) from mst_jobs j JOIN mst_companies c on j.company_id = c.id where c.name like ?"
	err := tx.QueryRow(SQL, "%"+companyName+"%").Scan(&result)
	helper.SendPanicError(err)

	return result
}
