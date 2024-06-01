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

	products, err := i.cartService.ListProducts(r.Context(), userID)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("list products error"), http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(products)
	if err != nil {
		cart.WriteResponse(logger, w, []byte("products error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(logger, w, buf, http.StatusOK)
}
