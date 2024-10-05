package entity

type Role struct {
	ID       string
	RoleName string
	Accounts []Account
}
