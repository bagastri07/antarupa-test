package model

type UserItem struct {
	ID     int64 `json:"id"`
	UserID int   `json:"userID"`
	ItemID int   `json:"itemID"`
	Total  int   `json:"total"`
}

type UserItemRepository interface {
	Create(userID, itemID, total int) error
	FindByUserIDAndItemID(userID, itemID int) (*UserItem, error)
	UpdateUserItemCounter(user_id, item_id int) error
}
