package repository

import (
	"context"
	"database/sql"
	"go-salaries-app/model/domain"
)

type SalaryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, salary domain.Salary) domain.Salary
	Update(ctx context.Context, tx *sql.Tx, salary domain.Salary) domain.Salary
	Delete(ctx context.Context, tx *sql.Tx, salary domain.Salary)
	FindById(ctx context.Context, tx *sql.Tx, salaryId int) (domain.Salary, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Salary
}
