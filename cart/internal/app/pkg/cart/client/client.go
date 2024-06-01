package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	errorapp "route256/cart/internal/app/pkg/errors"
	"route256/cart/internal/app/pkg/model"
	"time"
)

const token = "testtoken"

type ProductService struct {
}

func NewProductServiceClient() *ProductService {
	return &ProductService{}
}

type GetProductsRequest struct {
	Token string `json:"token"`
	Sku   uint32 `json:"sku"`
}

type GetProductsResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *ProductService) GetProductBySKU(ctx context.Context, sku model.SKU) (*model.Product, error) {
	data := &GetProductsRequest{
		Token: token,
		Sku:   uint32(sku),
	}

	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := "http://route256.pavl.uk:8080/get_product"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(marshalData))
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

	product := &model.Product{
		SKU:   sku,
		Name:  productRaw.Name,
		Price: productRaw.Price,
	}
	return product, nil
}
