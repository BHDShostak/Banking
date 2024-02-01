package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "<MAR>",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1990-01-01",
			Status:      "active",
			Email:       "<EMAIL>",
			Phone:       "123456789",
			Address:     "123 Main St",
			Password:    "<PASSWORD>",
		},
		{
			Id:          "2",
			Name:        "<BOH>",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1990-01-01",
			Status:      "active",
			Email:       "<EMAIL>",
			Phone:       "123456789",
			Address:     "123 Main St",
			Password:    "<PASSWORD>",
		},
		{
			Id:          "3",
			Name:        "<ANN>",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1990-01-01",
			Status:      "active",
			Email:       "<EMAIL>",
			Phone:       "123456789",
			Address:     "123 Main St",
			Password:    "<PASSWORD>",
		},
	}
	return CustomerRepositoryStub{customers}
}
