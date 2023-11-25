package dto

import (
	"banking/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (req NewAccountRequest) Validate() *errs.AppError {
	if req.Amount < 5000 {
		return errs.NewValidationError("To Open a new account you need to have 5000")
	}
	if strings.ToLower(req.AccountType) != "saving" &&
		strings.ToLower(req.AccountType) != "checking" {
		return errs.NewValidationError("Account type must be checking or saving")
	}
	return nil
}
