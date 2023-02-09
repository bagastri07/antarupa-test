package repository

import (
	"github.com/bagastri07/antarupa-test/model"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userCurrencyRepo struct {
	db *gorm.DB
}

func NewUserCurrencyRepo(db *gorm.DB) model.UserCurrencyRepository {
	return &userCurrencyRepo{
		db: db,
	}
}

func (r *userCurrencyRepo) GetBalanceByUserIDAndCurrencyType(userID, CurrencyType int) (int64, error) {
	balance := new(int64)
	err := r.db.Model(model.UserCurrency{}).
		Select("amount").
		Where("user_id = ?", userID).
		Where("currency_type = ?", CurrencyType).
		Scan(balance).
		Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return *balance, nil
}

func (r *userCurrencyRepo) UpdateBalance(userID, currencyType int, itemPrice int64) error {
	tx := r.db.Begin()

	userCurrency := new(model.UserCurrency)
	err := tx.First(userCurrency, "user_id = ? AND currency_type = ?", userID, currencyType).Error
	if err != nil {
		log.Error(err)
		return err
	}

	err = tx.Model(&model.UserCurrency{}).
		Where("user_id = ? AND currency_type = ?", userID, currencyType).
		Update("amount", userCurrency.Amount-itemPrice).
		Error

	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
