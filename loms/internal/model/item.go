package model

type Item struct {
	SKU   SKU
	Count uint16
}

type SKU int64

func (sku SKU) GetKey() int64 {
	return int64(sku)
}
