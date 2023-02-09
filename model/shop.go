package model

type Shop struct {
	ID           int    `json:"id"`
	ItemID       int    `json:"itemID"`
	ItemName     string `json:"itemName"`
	MaxOwned     int    `json:"maxOwned"`
	Price        int    `json:"price"`
	CurrencyType int    `json:"currencyType"`
	CurrencyName string `json:"currencyName"`
}

type ShopRepository interface {
	FindByItemID(itemID int) (*Shop, error)
}
