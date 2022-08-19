package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"go-salaries-app/app"
	"go-salaries-app/controller"
	"go-salaries-app/exception"
	"go-salaries-app/middleware"
	"go-salaries-app/repository"
	"go-salaries-app/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validator := validator.New()
	salaryRepository := repository.NewSalaryRepository()
	salaryService := service.NewSalaryService(salaryRepository, db, validator)
	salaryController := controller.NewSalaryController(salaryService)

	router := httprouter.New()

	router.GET("/api/salaries", salaryController.FindAll)
	router.GET("/api/salaries/:salaryId", salaryController.FindById)
	router.POST("/api/salaries", salaryController.Create)
	router.PUT("/api/salaries/:salaryId", salaryController.Update)
	router.DELETE("/api/salaries/:salaryId", salaryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleWare(router),
	}
	server.ListenAndServe()
}
