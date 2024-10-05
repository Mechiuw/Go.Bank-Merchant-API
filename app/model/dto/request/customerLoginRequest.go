package request

type CustomerLoginRequest struct {
	Email    string `json:"email"`    // Customer's email
	Password string `json:"password"` // Customer's password
}
