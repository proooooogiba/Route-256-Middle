package order

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *OrderRepo) GetOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	queryOrders := `SELECT id, user_id, status FROM orders WHERE id = $1`
	queryItemsOrders := `SELECT sku, item_count FROM items_orders WHERE order_id = $1`

	var order model.Order
	err := r.conn.QueryRow(ctx, queryOrders, id).Scan(
		&order.ID,
		&order.UserID,
		&order.Status,
	)
	if err != nil {
		return nil, errors.Wrap(err, "select order")
	}

	rows, err := r.conn.Query(ctx, queryItemsOrders, id)
	if err != nil {
		return nil, errors.Wrap(err, "select items_order")
	}
	defer rows.Close()

	items := make([]*model.Item, 0)
	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.SKU,
			&item.Count,
		)
		if err != nil {
			return nil, errors.Wrap(err, "rows.scan")
		}

		items = append(items, &item)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows.err")
	}

	order.Items = items

	return &order, nil
}
