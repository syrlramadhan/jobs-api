package dto

type CompanyCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
