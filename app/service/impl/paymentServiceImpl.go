package impl

import (
	"bank/app/model/dto/request"
	"bank/app/model/dto/response"
	"bank/app/model/entity"
	"bank/app/repository"
	"time"
)

type PaymentService struct {
	UserRepo    *repository.UserRepository
	HistoryRepo *repository.HistoryRepository
}

func (s *PaymentService) ProcessPayment(req request.PaymentRequest) (*response.PaymentResponse, error) {
	// Check if user is registered
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userFound bool
	for _, user := range users {
		if user.ID == req.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		return &response.PaymentResponse{IsSuccess: false, Message: "User not registered"}, nil
	}

	// Process payment
	transaction := entity.Transaction{
		ID:        "unique",
		UserID:    req.UserID,
		Amount:    req.Amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if err := s.HistoryRepo.SaveTransaction(transaction); err != nil {
		return nil, err
	}
	return &response.PaymentResponse{IsSuccess: true, Message: "Payment successful"}, nil
}
