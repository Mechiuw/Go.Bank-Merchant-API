package controller

import (
	"bank/app/model/dto/request"
	"bank/app/service/impl"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	AuthService *impl.AuthService
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req request.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userResp, err := c.AuthService.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(userResp)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// Handle logout
	w.Write([]byte("Logout successful"))
}
