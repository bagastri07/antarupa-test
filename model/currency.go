package model

type Currency struct {
	ID           int    `json:"id"`
	CurrencyName string `json:"currencyName"`
}

type CurrencyRepository interface{}
