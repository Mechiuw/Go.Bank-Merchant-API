package controller

import (
	"bank/app/model/dto/request"
	"bank/app/service/impl"
	"encoding/json"
	"net/http"
)

type PaymentController struct {
	PaymentService *impl.PaymentService
}

func (c *PaymentController) MakePayment(w http.ResponseWriter, r *http.Request) {
	var req request.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := c.PaymentService.ProcessPayment(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
