package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/pkg/shard_manager"
)

type shardManager interface {
	GetShardIndex(key shard_manager.ShardKey) shard_manager.ShardIndex
	GetShardIndexFromID(id int64) shard_manager.ShardIndex
	Pick(key shard_manager.ShardIndex) (*pgxpool.Pool, error)
}

type OrderRepo struct {
	sm shardManager
}

func NewOrderRepository(sm shardManager) *OrderRepo {
	return &OrderRepo{
		sm: sm,
	}
}
