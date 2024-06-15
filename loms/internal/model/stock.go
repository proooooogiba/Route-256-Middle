package model

type Stock struct {
	Sku        SKU    `json:"sku"`
	TotalCount uint64 `json:"total_count"`
	Reserved   uint64 `json:"reserved"`
}
