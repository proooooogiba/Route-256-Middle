package handlers

import (
	"net/http"
	"route256/cart/internal/app/cart"
	"route256/cart/internal/app/pkg/model"
	"strconv"
)

func (i *Implementation) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		cart.WriteResponse(w, []byte("parse user_id error"), http.StatusBadRequest)
		return
	}

	skuRaw := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(skuRaw, 10, 64)
	if err != nil {
		cart.WriteResponse(w, []byte("parse sku_id error"), http.StatusBadRequest)
		return
	}
	err = i.cartService.DeleteItem(r.Context(), userID, model.SKU(sku))
	if err != nil {
		cart.WriteResponse(w, []byte("delete item error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(w, []byte(`{}`), http.StatusNoContent)
	return
}
