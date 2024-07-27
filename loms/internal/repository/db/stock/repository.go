package stock

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type StockRepo struct {
	conn *pgxpool.Pool
}

func NewStockRepository(conn *pgxpool.Pool) *StockRepo {
	return &StockRepo{
		conn: conn,
	}
}
