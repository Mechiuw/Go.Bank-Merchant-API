package service

import (
	"bank/app/model/dto/request"
	"bank/app/model/dto/response"
	"bank/app/model/entity"
)

type PaymentService interface {
	CreatePayment(paymentRequest request.PaymentRequest) (response.PaymentResponse, error)
	GetAllPayment(page, pageSize int) ([]entity.Payment, error)
}
