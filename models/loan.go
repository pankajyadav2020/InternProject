package models

type Loan struct {
	ID           int    `json:"id" gorm:"primary_key"`
	CustomerName string `json:"customername"`
	PhoneNO      string `json:"phoneno"`
	Email        string `json:"email"`
	LoanAmount   int    `json:"loanamount"`
	CreditScore  int    `json:"creditscore"`
	Status       string `json:"status"`
}
