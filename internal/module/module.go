package module

type Products struct {
	ProductId    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductCount int     `json:"product_count"`
}

type JsonResponse struct {
	Data   Products `json:"data"`
	Source string   `json:"source"`
}
