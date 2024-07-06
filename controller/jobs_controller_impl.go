package controller

import (
	"geemod/jobs-api/dto"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/service"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type JobsControllerImpl struct{
	JobsService service.JobsService
	mu sync.Mutex
}

func NewJobsController(jobsService service.JobsService) JobsController{
	return &JobsControllerImpl{
		JobsService: jobsService,
	}
}

func (controller *JobsControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	
	// kunci end-point dengan golang mutex sebelum di akses orang lain ( race condition)
	controller.mu.Lock()
	defer controller.mu.Unlock()
	
	jobsRequest := dto.JobsCreateRequest{}
	helper.ReadFromRequestBody(request, &jobsRequest)

	jobsResponse := controller.JobsService.Create(request.Context(), jobsRequest)
	response := dto.Response{
		Code: 200,
		Status: "OK",
		Data: jobsResponse,
	}

	helper.WriteToResponseBody(writter, response)
	
}

func (controller *JobsControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	
	search := request.URL.Query().Get("search")
	limitStr := request.URL.Query().Get("limit")
	
	// default value for limit
	limit := helper.GetLimitValue(limitStr)

	jobs := controller.JobsService.FindAll(request.Context(), search, limit)
	count := controller.JobsService.Count(request.Context(), search)
	response := dto.ResponseList{
		Code: 200,
		Status: "OK",
		Data: jobs,
		Total: count,
		Limit: limit,
	}

	writter.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writter, response)
}