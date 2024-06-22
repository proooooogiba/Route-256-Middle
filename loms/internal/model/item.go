package model

type Item struct {
	SKU   SKU    `db:"sku"`
	Count uint16 `db:"item_count"`
}

type SKU int64

func (sku SKU) GetKey() int64 {
	return int64(sku)
}
