package repository

import (
	"bank/app/model/entity"
	"encoding/json"
	"io/ioutil"
)

type UserRepository struct {
	Filepath string
}

func (repo *UserRepository) LoadUsers() ([]entity.User, error) {
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
