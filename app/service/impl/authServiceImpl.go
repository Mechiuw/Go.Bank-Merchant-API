package impl

import (
	"bank/app/model/dto/request"
	"bank/app/model/dto/response"
	"bank/app/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Login(req request.UserRequest) (*response.UserResponse, error) {
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		// Compare the provided password with the stored hashed password
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err == nil && user.Username == req.Username {
			// Password is correct, return user response
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
	user, err := s.UserRepo.FindUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	user.IsLogged = false // Mark user as logged out
	return nil
}
