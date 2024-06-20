package http_handlers

import (
	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"net/http"
	"strconv"
)

func (i *Implementation) ClearCart(w http.ResponseWriter, r *http.Request) {
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

	err = i.cartService.Clear(r.Context(), userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			WriteErrorResponse(w, err, http.StatusNotFound)
			return
		}

		WriteErrorResponse(w, errors.Wrap(err, "clear cart error"), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, []byte(`{}`), http.StatusNoContent)
	return
}
