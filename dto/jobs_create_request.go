package dto

type JobsCreateRequest struct {
	CompanyId   string `validate:"required" json:"company_id"`
	Tittle      string `validate:"required" json:"tittle"`
	Description string `validate:"required" json:"description"`
}
