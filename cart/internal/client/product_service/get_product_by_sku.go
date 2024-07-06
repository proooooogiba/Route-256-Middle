package product_service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"io"
	"net/http"
	"net/url"
	"time"
)

const handlerGetProduct = "get_product"

func (c *ProductService) GetProductBySKU(ctx context.Context, sku model.SKU) (*model.Product, error) {
	c.limiterGetProduct.Take()
	c.mu.RLock()
	c.mu.RUnlock()

	data := &GetProductsRequest{
		Token: c.token,
		Sku:   uint32(sku),
	}

	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	path, err := url.JoinPath(c.basePath, handlerGetProduct)
	if err != nil {
		return nil, errors.Wrapf(err, "incorrect base basePath for %q", handlerGetProduct)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewBuffer(marshalData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var productRaw GetProductsResponse

	err = json.Unmarshal(body, &productRaw)
	if err != nil {
		return nil, err
	}

	if productRaw.Name == "" && productRaw.Price == 0 {
		return nil, errorapp.ErrNotFoundInPS
	}

	return getProduct(sku, productRaw), nil
}
