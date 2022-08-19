package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-salaries-app/helper"
	"go-salaries-app/model/domain"
)

type SalaryRepositoryImpl struct {
}

func NewSalaryRepository() SalaryRepository {
	return &SalaryRepositoryImpl{}
}

func (repository *SalaryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, salary domain.Salary) domain.Salary {
	sql := "insert into salary (role,company,expr,amount) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, salary.Role, salary.Company, salary.Expr, salary.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	salary.Id = int(id)
	return salary
}

func (repository *SalaryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, salary domain.Salary) domain.Salary {
	sql := "update salary set role=?, company=?, expr=?, amount=? where id_sal=?"
	_, err := tx.ExecContext(ctx, sql, salary.Role, salary.Company, salary.Expr, salary.Amount, salary.Id)
	helper.PanicIfError(err)

	return salary
}

func (repository *SalaryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, salary domain.Salary) {
	sql := "delete from salary where id_sal=?"
	_, err := tx.ExecContext(ctx, sql, salary.Id)
	helper.PanicIfError(err)
}

func (repository *SalaryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, salaryId int) (domain.Salary, error) {
	sql := "select * from salary where id_sal=?"
	rows, err := tx.QueryContext(ctx, sql, salaryId)
	defer rows.Close()
	helper.PanicIfError(err)

	salary := domain.Salary{}
	if rows.Next() {
		err := rows.Scan(&salary.Id, &salary.Role, &salary.Company, &salary.Expr, &salary.Amount)
		helper.PanicIfError(err)
		return salary, nil
	} else {
		return salary, errors.New("Salary is not found")
	}
}

func (repository *SalaryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Salary {
	sql := "select * from salary"
	var salaries []domain.Salary
	rows, err := tx.QueryContext(ctx, sql)
	defer rows.Close()
	
	for rows.Next() {
		salary := domain.Salary{}
		err := rows.Scan(&salary.Id, &salary.Role, &salary.Company, &salary.Expr, &salary.Amount)
		helper.PanicIfError(err)
		salaries = append(salaries, salary)
	}
	helper.PanicIfError(err)
	return salaries
}
