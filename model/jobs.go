package model

import "time"

type Jobs struct{
	Id string
	CompanyId string
	Tittle string
	Description string
	CreateAt time.Time
	Company *Company
}