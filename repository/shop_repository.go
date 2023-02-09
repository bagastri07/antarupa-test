package repository

import (
	"github.com/bagastri07/antarupa-test/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type shopRepo struct {
	db *gorm.DB
}

func NewShopRepo(db *gorm.DB) model.ShopRepository {
	return &shopRepo{
		db: db,
	}
}

func (r *shopRepo) FindByItemID(itemID int) (*model.Shop, error) {
	shop := new(model.Shop)
	query := r.db.
		Model(shop).
		Select("id, shops.item_id, max_owned, price, currency_type, items.item_name, currencies.currency_name").
		Joins("JOIN items ON shops.item_id = items.item_id").
		Joins("JOIN currencies ON shops.currency_type = currencies.currency_id").
		Where("shops.item_id = ?", itemID).
		Scan(shop)
	if query.RowsAffected <= 0 {
		return nil, nil
	}
	if err := query.Error; err != nil {
		logrus.Info(err)
		return nil, err
	}

	return shop, nil
}
