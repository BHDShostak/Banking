package service

import "banking/domain"

type Service interface {
	getAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) getAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewDefaultCustomerService(repo domain.CustomerRepository) Service {
	return DefaultCustomerService{repo}
}
