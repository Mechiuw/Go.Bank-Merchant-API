package service

import (
	"bank/app/model/dto/response"
	"bank/app/model/entity"
)

type Merchant interface {
	GetByIdMerchant(merchantID string) (response.MerchantResponse, error)
	GetAllMerchant(page, pageSize int) ([]entity.Merchant, error)
}
