package order

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type OrderRepo struct {
	conn *pgx.Conn
}

func NewOrderRepository(conn *pgx.Conn) *OrderRepo {
	return &OrderRepo{
		conn: conn,
	}
}
