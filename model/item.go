package model

type Item struct {
	ItemID   int    `json:"itemID"`
	ItemName string `json:"itemName"`
}

type PurchaseItemPayload struct {
	UserID int `json:"userID" validate:"required"`
	ItemID int `json:"itemID" validate:"required"`
}

type ItemRepository interface {
	FindByID(ID int) (*Item, error)
}
