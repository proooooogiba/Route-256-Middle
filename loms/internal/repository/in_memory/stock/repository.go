package stock

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

type StockInMemoryRepo struct {
	stocks map[model.SKU]*model.Stock
	mtx    sync.RWMutex
}

func NewStockRepository() (*StockInMemoryRepo, error) {
	jsonFile, err := os.Open("resources/stock-data.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stock-data.json")
	}
	defer func() {
		_ = jsonFile.Close()
	}()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read stock-data.json")
	}

	var initialData []model.Stock
	if err = json.Unmarshal(byteValue, &initialData); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal stock-data.json")
	}

	stocks := make(map[model.SKU]*model.Stock, len(initialData))

	for _, stock := range initialData {
		stocks[stock.Sku] = &stock
	}

	return &StockInMemoryRepo{
		stocks: stocks,
		mtx:    sync.RWMutex{},
	}, nil
}
