package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsLogged bool   `json:"isLogged"`
}
