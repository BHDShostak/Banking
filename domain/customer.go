package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
	Email       string
	Phone       string
	Address     string
	Password    string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
