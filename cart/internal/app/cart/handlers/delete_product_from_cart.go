package handlers

import (
	"log"
	"net/http"
	"route256/cart/internal/app/cart"
	"route256/cart/internal/app/pkg/middleware"
	"route256/cart/internal/app/pkg/model"
	"strconv"
)

func (i *Implementation) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
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

	skuRaw := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(skuRaw, 10, 64)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("parse sku_id error"), http.StatusBadRequest)
		return
	}
	err = i.cartService.DeleteItem(r.Context(), userID, model.SKU(sku))
	if err != nil {
		cart.WriteResponse(logger, w, []byte("delete item error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(logger, w, []byte(`{}`), http.StatusNoContent)
	return
}
