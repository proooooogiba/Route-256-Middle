package handlers

import (
	"github.com/pkg/errors"
	"net/http"
	"route256/cart/internal/app/cart"
	errorapp "route256/cart/internal/app/pkg/errors"
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

	if sku <= 0 || userID <= 0 {
		cart.WriteResponse(w, []byte("invalid params"), http.StatusBadRequest)
		return
	}

	err = i.cartService.DeleteItem(r.Context(), userID, model.SKU(sku))
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			cart.WriteResponse(w, []byte("invalid user"), http.StatusBadRequest)
			return
		}
		cart.WriteResponse(w, []byte("delete item error"), http.StatusInternalServerError)
		return
	}

	cart.WriteResponse(w, []byte(`{}`), http.StatusNoContent)
	return
}
