package model

type UserCurrency struct {
	UserCurrencyID int   `json:"userCurrencyID"`
	UserID         int   `json:"userID"`
	CurrencyType   int   `json:"currencyType"`
	Amount         int64 `json:"amount"`
}

type UserCurrencyRepository interface {
	GetBalanceByUserIDAndCurrencyType(userID, currencyType int) (int64, error)
	UpdateBalance(userID, currencyType int, itemPrice int64) error
}

func (UserCurrency) TableName() string {
	return "user_currency"
}
