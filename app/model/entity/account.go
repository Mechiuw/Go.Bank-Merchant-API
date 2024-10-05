package entity

import "time"

type Account struct {
	Email     string
	Password  string
	LoginTime time.Time
	RoleID    string
	UserID    string
}
