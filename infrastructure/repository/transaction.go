package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/valentimcanejo/codepix/domain/model"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (t TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.Db.Preload("AccountForm.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction found")
	}

	return &transaction, nil
}

func (t TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}

	return nil
}
