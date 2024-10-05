package repository

import (
	"bank/app/model/entity"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type UserRepository struct {
	Filepath string
}

func (repo *UserRepository) GetAllUsers() ([]entity.User, error) {
	file, err := ioutil.ReadFile(repo.Filepath)
	if err != nil {
		return nil, err
	}
	var users []entity.User
	err = json.Unmarshal(file, &users)
	return users, err
}

func (repo *UserRepository) SaveUsers(users []entity.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(repo.Filepath, data, 0644)
}

func (r *UserRepository) FindUser(username string) (entity.User, error) {
	users, err := r.GetAllUsers()
	if err != nil {
		return entity.User{}, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (r *UserRepository) FindUserByID(userID string) (entity.User, error) {
	users, err := r.GetAllUsers()
	if err != nil {
		return entity.User{}, err
	}

	for _, user := range users {
		if user.ID == userID {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}
