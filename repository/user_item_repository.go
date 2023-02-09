package repository

import (
	"errors"
	"time"

	"github.com/bagastri07/antarupa-test/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userItemRepo struct {
	db *gorm.DB
}

func NewUserItemRepo(db *gorm.DB) model.UserItemRepository {
	return &userItemRepo{
		db: db,
	}
}

func (r *userItemRepo) Create(userID, itemID, total int) error {
	userItem := model.UserItem{
		ID:     time.Now().UnixMilli(),
		UserID: userID,
		ItemID: itemID,
		Total:  total,
	}

	tx := r.db.Begin()
	err := tx.Create(&userItem).Error
	if err != nil {
		logrus.Error(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *userItemRepo) FindByUserIDAndItemID(userID, itemID int) (*model.UserItem, error) {
	userItem := new(model.UserItem)

	err := r.db.First(userItem, "user_id = ? AND item_id = ?", userID, itemID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}

	return userItem, nil
}

func (r *userItemRepo) UpdateUserItemCounter(user_id, item_id int) error {
	tx := r.db.Begin()
	userItem := new(model.UserItem)
	err := tx.First(userItem, "user_id = ? AND item_id = ?", user_id, item_id).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = tx.Model(&model.UserItem{}).
		Where("user_id = ?", user_id).
		Where("item_id = ?", item_id).
		Update("total", userItem.Total+1).
		Error

	if err != nil {
		logrus.Error(err)
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
