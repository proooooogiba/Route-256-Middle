package http_handlers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
	"strconv"
)

type ListCartProductsResponse struct {
	Items      []*Item `json:"items"`
	TotalPrice uint32  `json:"total_price"`
}

type Item struct {
	SKU   model.SKU `json:"sku_id"`
	Name  string    `json:"name"`
	Count uint16    `json:"count"`
	Price uint32    `json:"price"`
}

func (i *Implementation) ListCartProducts(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse sku_id error"), http.StatusBadRequest)
		return
	}

	if userID <= 0 {
		WriteErrorResponse(w, errorapp.ErrInvalidUserId, http.StatusBadRequest)
		return
	}

	products, err := i.cartService.ListProducts(r.Context(), userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			WriteErrorResponse(w, errorapp.ErrNotFoundUser, http.StatusNotFound)
			return
		}
		if errors.Is(err, errorapp.ErrOutOfStock) {
			WriteErrorResponse(w, errorapp.ErrOutOfStock, http.StatusPreconditionFailed)
			return
		}
		WriteErrorResponse(w, errors.Wrap(err, "list products error"), http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(products)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "products marshal error"), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, buf, http.StatusOK)
}
