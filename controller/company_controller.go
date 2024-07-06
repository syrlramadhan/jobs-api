package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CompanyController interface{
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}