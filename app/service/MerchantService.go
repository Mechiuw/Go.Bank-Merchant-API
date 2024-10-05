package service

import (
	"bank/app/model/dto/request"
	"bank/app/model/dto/response"
	"bank/app/model/entity"
)

type Merchant interface {
	CreateMerchant(request request.MerchantRequest) (response.MerchantResponse, error)
	UpdateMerchant(merchantID string, request request.MerchantRequest) (response.MerchantResponse, error)
	GetByIdMerchant(merchantID string) (response.MerchantResponse, error)
	GetAllMerchant(page, pageSize int) ([]entity.Merchant, error)
}
