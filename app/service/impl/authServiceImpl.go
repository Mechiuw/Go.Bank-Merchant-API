package impl

import (
	"bank/app/model/dto/request"
	"bank/app/model/dto/response"
	"bank/app/repository"
	"errors"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Login(req request.UserRequest) (*response.UserResponse, error) {
	users, err := s.UserRepo.LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == req.Username && user.Password == req.Password {
			return &response.UserResponse{
				ID:       user.ID,
				Username: user.Username,
				Role:     user.Role,
			}, nil
		}
	}
	return nil, errors.New("invalid credentials")
}

func (s *AuthService) Logout(userID string) error {
	// Implement logout logic
	return nil
}
