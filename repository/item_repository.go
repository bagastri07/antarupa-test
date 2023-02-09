package repository

import (
	"errors"

	"github.com/bagastri07/antarupa-test/model"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) model.ItemRepository {
	return &itemRepo{
		db: db,
	}
}

func (r *itemRepo) FindByID(ID int) (*model.Item, error) {
	item := new(model.Item)
	err := r.db.First(item, "item_id = ?", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}

	return item, nil
}
