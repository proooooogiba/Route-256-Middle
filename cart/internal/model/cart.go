package model

type Cart struct {
	Items []Item
}

type Item struct {
	SKU   SKU
	Count uint16
}
