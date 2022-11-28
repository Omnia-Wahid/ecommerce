package model

type ProductStock struct {
	Country        string `json:"country,omitempty"`
	Sku            string `json:"suk"`
	Name           string `json:"name,omitempty"`
	StockChange    int    `json:"stock_change,omitempty"`
	ConsumedAmount int    `json:"consumed_amount,omitempty"`
}
