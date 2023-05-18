package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	UserId    uint    `gorm:"not null;" json:"user_id"`
	Name      string  `gorm:"size:255; not null;" json:"name"`
	Reference string  `gorm:"size:255; null;" json:"reference"`
	Amount    float64 `gorm:"not null;" json:"amount"`
}

func (t *Transaction) SaveTransaction() (*Transaction, error) {
	err := DB.Create(&t).Error
	if err != nil {
		return &Transaction{}, err
	}

	return t, nil
}

func GetTransactionByID(uid, tid uint) (t Transaction, err error) {
	if err = DB.Where("user_id = ?", uid).First(&t, tid).Error; err != nil {
		return t, err
	}

	return t, nil
}

func GetTransactions(uid uint) (t []Transaction, err error) {
	if err = DB.Where("user_id = ?", uid).Find(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}
