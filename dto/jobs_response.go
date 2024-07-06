package dto

import "time"

type JobsResponse struct{
	Id string `json:"id"`
	CompanyId string `json:"company_id"`
	CompanyName string `json:"company_name"`
	Tittle string `json:"tittle"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}