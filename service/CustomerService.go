// PRIMARY PORT ->
package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

// PORT ->
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// DEPENDENCY INJECTION ?
type DefaultCustomerService struct {
	r domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.r.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.r.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

// HELPER FUNC
func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{r}
}
