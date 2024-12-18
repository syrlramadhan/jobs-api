package main

import (
	"fmt"
	"geemod/jobs-api/config"
	"geemod/jobs-api/controller"
	"geemod/jobs-api/exception"
	"geemod/jobs-api/helper"
	"geemod/jobs-api/repository"
	"geemod/jobs-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("--------- JOBS API ---------")

	db := config.ConnectDB()
	validate := validator.New()

	// create object for company
	companyRespository := repository.NewCompanyRepository()
	companyService := service.NewCompanyService(companyRespository, db, validate)
	companyController := controller.NewCompanyController(companyService)

	// create object for jobs
	jobsRepository := repository.NewJobsRepository()
	jobsService := service.NewJobsService(jobsRepository, companyRespository, db, validate)
	jobsController := controller.NewJobsController(jobsService)

	router := httprouter.New()

	// end-point api for companies
	router.POST("/api/companies", companyController.Create)
	router.PUT("/api/companies/:companyId", companyController.Update)

	// end-point api for jobs
	router.POST("/api/jobs", jobsController.Create)
	router.GET("/api/jobs", jobsController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.SendPanicError(err)
}
