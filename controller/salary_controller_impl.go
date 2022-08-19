package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-salaries-app/helper"
	"go-salaries-app/model/web"
	"go-salaries-app/service"
	"net/http"
	"strconv"
)

type SalaryControllerImpl struct {
	service.SalaryService
}

func NewSalaryController(salaryService service.SalaryService) SalaryController {
	return &SalaryControllerImpl{SalaryService: salaryService}
}

func (controller *SalaryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.SalaryCreateRequest{}
	helper.ReadFromRequestBody(r, &request)

	create := controller.SalaryService.Create(r.Context(), request)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   create,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *SalaryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.SalaryUpdateRequest{}
	helper.ReadFromRequestBody(r, &request)

	id := params.ByName("salaryId")
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	request.Id = idInt

	create := controller.SalaryService.Update(r.Context(), request)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   create,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *SalaryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("salaryId")
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.SalaryService.Delete(r.Context(), idInt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Salary has been deleted",
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *SalaryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("salaryId")
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	findById := controller.SalaryService.FindById(r.Context(), idInt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findById,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *SalaryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	findAll := controller.SalaryService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findAll,
	}

	helper.WriteToResponseBody(w, &webResponse)
}
