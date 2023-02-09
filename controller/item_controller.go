package controller

import (
	"net/http"

	cErr "github.com/bagastri07/antarupa-test/custom_error"
	"github.com/bagastri07/antarupa-test/model"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ItemController struct {
	userRepo         model.UserRepository
	shopRepo         model.ShopRepository
	userCurrencyRepo model.UserCurrencyRepository
	userItemRepo     model.UserItemRepository
}

func NewItemController(userRepo model.UserRepository, shopRepo model.ShopRepository, userCurrencyRepo model.UserCurrencyRepository, userItemRepo model.UserItemRepository) *ItemController {
	return &ItemController{
		userRepo:         userRepo,
		shopRepo:         shopRepo,
		userCurrencyRepo: userCurrencyRepo,
		userItemRepo:     userItemRepo,
	}
}

func (ctrl *ItemController) PurchaseItem(c echo.Context) error {
	purchaseItemPayload := new(model.PurchaseItemPayload)

	if err := c.Bind(purchaseItemPayload); err != nil {
		log.Error(err)
		return cErr.ErrInternalServerErr
	}

	if err := c.Validate(purchaseItemPayload); err != nil {
		log.Error(err)
		return err
	}

	if err := ctrl.validateUser(purchaseItemPayload.UserID); err != nil {
		return err
	}

	shopItem, err := ctrl.getShopItem(purchaseItemPayload.ItemID)
	if err != nil {
		return err
	}

	userBalance, err := ctrl.userCurrencyRepo.GetBalanceByUserIDAndCurrencyType(purchaseItemPayload.UserID, shopItem.CurrencyType)
	if err != nil {
		log.Error(err)
		return cErr.ErrInternalServerErr
	}

	if userBalance < int64(shopItem.Price) {
		return cErr.ErrNotEnoughGameCurrency
	}

	userItem, err := ctrl.userItemRepo.FindByUserIDAndItemID(purchaseItemPayload.UserID, shopItem.ItemID)
	if err != nil {
		log.Error(err)
		return cErr.ErrInternalServerErr
	}

	if userItem == nil {
		if err := ctrl.userItemRepo.Create(purchaseItemPayload.UserID, shopItem.ItemID, 1); err != nil {
			log.Error(err)
			return cErr.ErrInternalServerErr
		}
	} else if userItem.Total+1 > shopItem.MaxOwned {
		return cErr.ErrMaxItemReached
	} else {
		err := ctrl.userItemRepo.UpdateUserItemCounter(purchaseItemPayload.UserID, shopItem.ItemID)
		if err != nil {
			log.Error(err)
			return cErr.ErrInternalServerErr
		}
	}

	err = ctrl.userCurrencyRepo.UpdateBalance(purchaseItemPayload.UserID, shopItem.CurrencyType, int64(shopItem.Price))
	if err != nil {
		log.Error(err)
		return cErr.ErrInternalServerErr
	}

	return c.JSON(http.StatusOK, model.MessageResponse{
		Message: "success",
	})
}

func (ctrl *ItemController) validateUser(userID int) error {
	user, err := ctrl.userRepo.FindByID(userID)
	if err != nil {
		log.Error(err)
		return cErr.ErrInternalServerErr
	}

	if user == nil {
		return cErr.NewHttpCustomErr(http.StatusNotFound, "user not found")
	}
	return nil
}

func (ctrl *ItemController) getShopItem(itemID int) (*model.Shop, error) {
	shopItem, err := ctrl.shopRepo.FindByItemID(itemID)
	if err != nil {
		log.Error(err)
		return nil, cErr.ErrInternalServerErr
	}

	if shopItem == nil {
		return nil, cErr.NewHttpCustomErr(http.StatusNotFound, "item not listed in the shop")
	}
	return shopItem, nil
}
