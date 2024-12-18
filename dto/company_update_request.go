package dto

type CompanyUpdateRequest struct {
	Id   string
	Name string `validate:"required" json:"name"`
}
