package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer{
		{Id: "0001", Name: "Deepak", City: "New Delhi", Zipcode: "110075", Status: "Active"},
		{Id: "0001", Name: "Chinmay", City: "New Delhi", Zipcode: "110075", Status: "Active"},
	}
	return CustomerRepositoryStub{customer}
}
