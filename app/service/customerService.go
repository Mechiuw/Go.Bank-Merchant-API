package service

import (
	"bank/app/model/dto/response"
	"bank/app/model/entity"
)

type CustomerService interface {
	GetAllCustomer(page, pageSize int) ([]entity.Customer, error)
	GetByIdCustomer(CustomerID string) (response.CustomerLoginResponse, error)
}
