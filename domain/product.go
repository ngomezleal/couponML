package domain

type Response struct {
	SiteId  string    `json:"site_id"`
	Seller  Seller    `json:"seller"`
	Results []Product `json:"results"`
}

type Product struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type ProductRepository interface {
	FindProductsByClientId(key string) (*Response, error)
	GetProductsByCouponAndClientId(key string, coupon float64) ([]Product, error)
}
