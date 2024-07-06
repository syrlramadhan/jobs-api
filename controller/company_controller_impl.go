package controller

import (
	"geemod/jobs-api/dto"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/service"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type CompanyControllerImpl struct{
	CompanyService service.CompanyService
	mu sync.Mutex
}

func NewCompanyController(companyService service.CompanyService) CompanyController{
	return &CompanyControllerImpl{
		CompanyService: companyService,
	}
}

func (controller *CompanyControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	
	// kunci end-point dengan golang mutex sebelum di akses orang lain ( race condition)
	controller.mu.Lock()
	defer controller.mu.Unlock()

	companyRequest := dto.CompanyCreateRequest{}
	helper.ReadFromRequestBody(request, &companyRequest)

	companyResponse := controller.CompanyService.Create(request.Context(), companyRequest)
	response := dto.Response{
		Code: 200,
		Status: "OK",
		Data: companyResponse,
	}

	helper.WriteToResponseBody(writter, response)

}

func (controller *CompanyControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	
	// kunci end-point dengan golang mutex sebelum di akses orang lain ( race condition)
	controller.mu.Lock()
	defer controller.mu.Unlock()
	
	companyRequest := dto.CompanyUpdateRequest{}
	helper.ReadFromRequestBody(request, &companyRequest)

	companyId := params.ByName("companyId")
	companyRequest.Id = companyId

	companyResponse := controller.CompanyService.Update(request.Context(), companyRequest)
	response := dto.Response{
		Code: 200,
		Status: "OK",
		Data: companyResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writter, response)
	
}
