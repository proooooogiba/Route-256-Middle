package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"route256/cart/internal/app/cart"
	"route256/cart/internal/app/pkg/middleware"
	"route256/cart/internal/app/pkg/model"
	"strconv"
)

type AddItemRequestBody struct {
	Count uint16 `json:"count"`
}

func (i *Implementation) AddItemToCart(w http.ResponseWriter, r *http.Request) {
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

	var body AddItemRequestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("parse body error"), http.StatusBadRequest)
		return
	}

	err = i.cartService.AddItem(r.Context(), userID, model.SKU(sku), body.Count)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("add item error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(logger, w, []byte(`{}`), http.StatusCreated)
	return
}
