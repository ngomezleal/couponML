package domain

type Product struct {
	Id       string  `json:"id"`
	SiteId   string  `json:"site_id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	SellerId int64   `json:"seller_id"`
}

type ProductRepository interface {
	FindTopProducts() ([]*Product, error)
	CalculateAndSaveProductsBought(input InputParams) ([]Product, error)
}
