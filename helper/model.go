package helper

import (
	"go-salaries-app/model/domain"
	"go-salaries-app/model/web"
)

func ToSalaryResponse(salary domain.Salary) web.SalaryResponse {
	return web.SalaryResponse{
		Id:      salary.Id,
		Role:    salary.Role,
		Company: salary.Company,
		Expr:    salary.Expr,
		Amount:  salary.Amount,
	}
}
