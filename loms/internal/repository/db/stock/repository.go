package stock

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type StockRepo struct {
	conn *pgx.Conn
}

func NewStockRepository(conn *pgx.Conn) *StockRepo {
	return &StockRepo{
		conn: conn,
	}
}
