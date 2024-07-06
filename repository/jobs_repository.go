package repository

import (
	"context"
	"database/sql"
	"geemod/jobs-api/model"
)

type JobsRepository interface{
	Save(ctx context.Context, tx *sql.Tx, jobs model.Jobs)model.Jobs
	FindAll(ctx context.Context, tx *sql.Tx) []model.Jobs
}