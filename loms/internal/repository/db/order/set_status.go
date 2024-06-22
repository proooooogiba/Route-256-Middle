package order

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *OrderRepo) SetStatus(ctx context.Context, id int64, status model.Status) error {
	_, err := r.conn.Exec(ctx, `UPDATE orders SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return errors.Wrap(err, "conn.Exec")
	}
	return nil
}
