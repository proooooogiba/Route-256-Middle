package order

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *OrderRepo) CreateOrder(ctx context.Context, userID int64, items []*model.Item) (*model.Order, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	queryOrders := `INSERT INTO orders(user_id, status) VALUES($1,$2) RETURNING id`
	var orderID int64
	err = tx.QueryRow(ctx, queryOrders, userID, model.New.String()).Scan(&orderID)
	if err != nil {
		return nil, errors.Wrap(err, "create order")
	}

	queryItemsOrders := `INSERT INTO items_orders(sku, order_id, item_count) VALUES($1,$2,$3)`
	for _, item := range items {
		_, err = tx.Exec(ctx, queryItemsOrders, item.SKU, orderID, item.Count)
		if err != nil {
			return nil, errors.Wrap(err, "add items to order")
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "commit transaction")
	}

	return &model.Order{
		ID:     orderID,
		UserID: userID,
		Status: model.New.String(),
		Items:  items,
	}, nil
}
