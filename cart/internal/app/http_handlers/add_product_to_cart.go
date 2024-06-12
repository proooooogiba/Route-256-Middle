package http_handlers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
	"strconv"
)

type AddItemRequestBody struct {
	Count uint16 `json:"count"`
}

func (i *Implementation) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse int"), http.StatusBadRequest)
		return
	}

	skuRaw := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(skuRaw, 10, 64)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse sku_id error"), http.StatusBadRequest)
		return
	}

	var body AddItemRequestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse body error"), http.StatusBadRequest)
		return
	}

	if sku <= 0 || userID <= 0 {
		WriteErrorResponse(w, errorapp.ErrInvalidParams, http.StatusBadRequest)
		return
	}
	if body.Count <= 0 {
		WriteErrorResponse(w, errorapp.ErrInvalidBody, http.StatusBadRequest)
		return
	}

	err = i.cartService.AddItem(r.Context(), userID, model.SKU(sku), body.Count)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundInPS) {
			WriteErrorResponse(w, err, http.StatusPreconditionFailed)
			return
		}

		WriteErrorResponse(w, errors.Wrap(err, "add item error"), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, []byte(`{}`), http.StatusOK)
	return
}
