package model

import "time"

type UserData struct {
	UserID   int       `json:"userID"`
	Username string    `json:"username"`
	JoinDate time.Time `json:"joinDate"`
}

type UserRepository interface {
	FindByID(ID int) (*UserData, error)
}
