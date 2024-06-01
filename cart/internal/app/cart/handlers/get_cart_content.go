package handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/app/cart"
	"route256/cart/internal/app/pkg/model"
	"strconv"
)

type ListCartProductsResponse struct {
	Items      []*Item `json:"items"`
	TotalPrice uint32  `json:"total_price"`
}

type Item struct {
	SKU   model.SKU `json:"sku_id"`
	Name  string    `json:"name"`
	Count uint16    `json:"cost"`
	Price uint32    `json:"price"`
}

func (i *Implementation) ListCartProducts(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		cart.WriteResponse(w, []byte("parse user_id error"), http.StatusBadRequest)
		return
	}

	products, err := i.cartService.ListProducts(r.Context(), userID)
	if err != nil {
		cart.WriteResponse(w, []byte("list products error"), http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(products)
	if err != nil {
		cart.WriteResponse(w, []byte("products error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(w, buf, http.StatusOK)
}
