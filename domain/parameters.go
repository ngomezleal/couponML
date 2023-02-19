package domain

type InputParams struct {
	Key string `json:"key"`
}

type InputParamsCoupon struct {
	Key    string  `json:"key"`
	Coupon float64 `json:"coupon"`
}
