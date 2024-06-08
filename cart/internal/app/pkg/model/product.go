package model

type Product struct {
	SKU   SKU    `json:"sku_id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type SKU int64

func (sku SKU) GetKey() int64 {
	return int64(sku)
}
