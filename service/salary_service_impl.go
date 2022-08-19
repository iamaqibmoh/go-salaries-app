package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-salaries-app/exception"
	"go-salaries-app/helper"
	"go-salaries-app/model/domain"
	"go-salaries-app/model/web"
	"go-salaries-app/repository"
)

type SalaryServiceImpl struct {
	repository.SalaryRepository
	*sql.DB
	*validator.Validate
}

func NewSalaryService(salaryRepository repository.SalaryRepository, DB *sql.DB, validate *validator.Validate) SalaryService {
	return &SalaryServiceImpl{SalaryRepository: salaryRepository, DB: DB, Validate: validate}
}

func (service *SalaryServiceImpl) Create(ctx context.Context, request web.SalaryCreateRequest) web.SalaryResponse {
	err := service.Validate.Struct(&request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	salary := domain.Salary{
		Role:    request.Role,
		Company: request.Company,
		Expr:    request.Expr,
		Amount:  request.Amount,
	}

	salary = service.SalaryRepository.Create(ctx, tx, salary)

	return helper.ToSalaryResponse(salary)
}

func (service *SalaryServiceImpl) Update(ctx context.Context, request web.SalaryUpdateRequest) web.SalaryResponse {
	err := service.Validate.Struct(&request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	salary, err := service.SalaryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	salary.Role = request.Role
	salary.Company = request.Company
	salary.Expr = request.Expr
	salary.Amount = request.Amount

	update := service.SalaryRepository.Update(ctx, tx, salary)

	return helper.ToSalaryResponse(update)
}

func (service *SalaryServiceImpl) Delete(ctx context.Context, salaryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	salary, err := service.SalaryRepository.FindById(ctx, tx, salaryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SalaryRepository.Delete(ctx, tx, salary)
}

func (service *SalaryServiceImpl) FindById(ctx context.Context, salaryId int) web.SalaryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	salary, err := service.SalaryRepository.FindById(ctx, tx, salaryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSalaryResponse(salary)
}

func (service *SalaryServiceImpl) FindAll(ctx context.Context) []web.SalaryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	salaries := service.SalaryRepository.FindAll(ctx, tx)
	var salaryResponses []web.SalaryResponse
	for _, salary := range salaries {
		salaryResponses = append(salaryResponses, helper.ToSalaryResponse(salary))
	}
	return salaryResponses
}
