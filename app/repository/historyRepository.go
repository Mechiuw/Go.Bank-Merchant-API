package repository

import (
	"bank/app/model/entity"
	"encoding/json"
	"io/ioutil"
)

type HistoryRepository struct {
	FilePath string
}

func (repo *HistoryRepository) LoadHistory() ([]entity.Transaction, error) {
	file, err := ioutil.ReadFile(repo.FilePath)
	if err != nil {
		return nil, err
	}
	var history []entity.Transaction
	err = json.Unmarshal(file, &history)
	return history, err
}

func (repo *HistoryRepository) SaveTransaction(transaction entity.Transaction) error {
	history, err := repo.LoadHistory()
	if err != nil {
		return err
	}
	history = append(history, transaction)
	return repo.SaveHistory(history)
}

func (repo *HistoryRepository) SaveHistory(history []entity.Transaction) error {
	data, err := json.Marshal(history)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(repo.FilePath, data, 0644)
}
