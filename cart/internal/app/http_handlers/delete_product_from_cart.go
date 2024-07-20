package http_handlers

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

func (i *Implementation) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse user_id error"), http.StatusBadRequest)
		return
	}

	skuRaw := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(skuRaw, 10, 64)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse sku_id error"), http.StatusBadRequest)
		return
	}

	if sku <= 0 || userID <= 0 {
		WriteErrorResponse(w, errorapp.ErrInvalidParams, http.StatusBadRequest)
		return
	}

	err = i.cartService.DeleteItem(r.Context(), userID, model.SKU(sku))
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			WriteErrorResponse(w, errorapp.ErrNotFoundUser, http.StatusNotFound)
			return
		}
		WriteErrorResponse(w, errors.Wrap(err, "delete item error"), http.StatusNotFound)
		return
	}

	WriteResponse(w, []byte(`{}`), http.StatusNoContent)
	return
}
