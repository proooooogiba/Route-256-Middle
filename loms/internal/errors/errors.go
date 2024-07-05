package errorapp

import "github.com/pkg/errors"

var (
	ErrOutOfStock    = errors.New("can't reserve items")
	ErrUnknownStatus = errors.New("unknown status")
	ErrOrderNotFound = errors.New("order not found")
	ErrStockNotFound = errors.New("stock not found")
	ErrSkuNotFound   = errors.New("sku not found")
	ErrNoNil         = errors.New("err no nil")
)
