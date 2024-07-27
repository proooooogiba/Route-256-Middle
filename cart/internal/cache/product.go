package cache

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"sync"
)

type Cacher interface {
	Get(ctx context.Context, sku model.SKU) (*model.Product, error)
	Set(ctx context.Context, product *model.Product) error
}

type key struct {
	sku model.SKU
}

// Buffer - реализация кольцевого буфера, согласно описанию
// из doc/homework-8
type Buffer struct {
	size   int64
	index  int64
	buffer []key
	data   *sync.Map
	mx     sync.Mutex
}

func New(size uint) *Buffer {
	if size == 0 {
		size = 100
	}

	buffer := make([]key, size)
	data := &sync.Map{}

	return &Buffer{
		size:   int64(size),
		index:  int64(size),
		buffer: buffer,
		data:   data,
	}
}

func (b *Buffer) Get(ctx context.Context, sku model.SKU) (*model.Product, error) {
	product, ok := b.data.Load(sku)
	if !ok {
		return nil, errors.New("sku not found")
	}

	return product.(*model.Product), nil
}

func (b *Buffer) Set(ctx context.Context, product *model.Product) error {
	if b.size == 0 {
		return errors.New("buffer max size is zero")
	}

	b.mx.Lock()
	oldSku := b.buffer[b.index%b.size].sku
	b.data.Delete(oldSku)
	b.buffer[b.index%b.size].sku = product.SKU
	b.index++
	b.mx.Unlock()

	b.data.Store(product.SKU, product)

	return nil
}
