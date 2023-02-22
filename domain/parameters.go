package domain

type InputParams struct {
	Items        []string `json:"item_ids"`
	CouponAmount float64  `json:"amount"`
}
