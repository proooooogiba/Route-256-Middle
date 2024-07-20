package http_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
)

type CheckoutResponse struct {
	OrderID int64 `json:"order_id"`
}

func (i *Implementation) Checkout(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.PathValue("user_id")
	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "parse user_id error"), http.StatusBadRequest)
		return
	}

	if userID <= 0 {
		WriteErrorResponse(w, errorapp.ErrInvalidUserId, http.StatusBadRequest)
		return
	}

	ckecoutResponse, err := i.cartService.Checkout(r.Context(), userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			WriteErrorResponse(w, err, http.StatusNotFound)
			return
		}

		WriteErrorResponse(w, errors.Wrap(err, "cart checkout error"), http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(ckecoutResponse)
	if err != nil {
		WriteErrorResponse(w, errors.Wrap(err, "products marshal error"), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, buf, http.StatusNoContent)
	return
}
