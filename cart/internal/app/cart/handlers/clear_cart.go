package handlers

import (
	"net/http"
	"route256/cart/internal/app/cart"
	"strconv"
)

func (i *Implementation) ClearCart(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		cart.WriteResponse(w, []byte("parse user_id error"), http.StatusBadRequest)
		return
	}

	err = i.cartService.Clear(r.Context(), userID)
	if err != nil {
		cart.WriteResponse(w, []byte("clear cart error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(w, []byte(`{}`), http.StatusNoContent)
	return
}
