package order

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *OrderRepo) SetStatus(ctx context.Context, id int64, status model.Status) error {
	shIndex := r.sm.GetShardIndexFromID(id)
	db, err := r.sm.Pick(shIndex)
	if err != nil {
		return errors.Wrap(err, "failed to pick db")
	}

	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err = db.Exec(ctx, query, status, id)
	if err != nil {
		return errors.Wrap(err, "conn.Exec")
	}
	return nil
}
