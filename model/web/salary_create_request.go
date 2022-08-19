package web

type SalaryCreateRequest struct {
	Role    string `validate:"required,min=1,max=200" json:"role"`
	Company string `validate:"required,min=1,max=200" json:"company"`
	Expr    int    `validate:"required,min=1,max=200" json:"expr"`
	Amount  int    `validate:"required,min=1" json:"amount"`
}
