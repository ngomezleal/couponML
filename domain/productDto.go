package domain

type OutputProductDto struct {
	Items []ProductDto `json:"item_ids"`
	Total float64      `json:"total"`
}

type ProductDto struct {
	Id       string  `json:"id"`
	SiteId   string  `json:"site_id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	SellerId int64   `json:"seller_id"`
}
