package domain

import "banking/errs"

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
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
