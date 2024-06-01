package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"route256/cart/internal/app/pkg/model"
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
	Name  string `json:"token"`
	Price uint32 `json:"sku"`
}

func (c *ProductService) GetProductBySKU(ctx context.Context, sku model.SKU) (*model.Product, error) {
	url := "http://route256.pavl.uk:8080/get_product"

	data := &GetProductsRequest{
		Token: token,
		Sku:   uint32(sku),
	}

	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshalData))
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

	product := &model.Product{
		SKU:   sku,
		Name:  productRaw.Name,
		Price: productRaw.Price,
	}
	return product, nil
}
