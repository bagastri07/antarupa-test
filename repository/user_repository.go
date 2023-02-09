package repository

import (
	"errors"

	"github.com/bagastri07/antarupa-test/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) model.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) FindByID(ID int) (*model.UserData, error) {
	user := new(model.UserData)

	err := r.db.First(user, "user_id = ?", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}

	return user, nil
}
