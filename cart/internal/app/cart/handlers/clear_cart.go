package handlers

import (
	"log"
	"net/http"
	"route256/cart/internal/app/cart"
	"route256/cart/internal/app/pkg/middleware"
	"strconv"
)

func (i *Implementation) ClearCart(w http.ResponseWriter, r *http.Request) {
	logger, err := middleware.GetLoggerFromContext(r.Context())
	if err != nil {
		log.Printf("can not get logger from context: %s", err)
		middleware.WriteNoLoggerResponse(w)
	}

	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("parse user_id error"), http.StatusBadRequest)
		return
	}

	err = i.cartService.Clear(r.Context(), userID)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("clear cart error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(logger, w, []byte(`{}`), http.StatusNoContent)
	return
}
