package handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/app/cart"
	"route256/cart/internal/app/pkg/model"
	"strconv"
)

type AddItemRequestBody struct {
	Count uint16 `json:"count"`
}

func (i *Implementation) AddItemToCart(w http.ResponseWriter, r *http.Request) {
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

	var body AddItemRequestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		cart.WriteResponse(w, []byte("parse body error"), http.StatusBadRequest)
		return
	}

	err = i.cartService.AddItem(r.Context(), userID, model.SKU(sku), body.Count)
	if err != nil {
		cart.WriteResponse(w, []byte("add item error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(w, []byte(`{}`), http.StatusCreated)
	return
}
