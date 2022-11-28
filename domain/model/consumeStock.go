package model

type ConsumeStock struct {
	Country        string `json:"country" validate:"required"`
	ConsumedAmount int    `json:"consumed_amount" validate:"required"`
}
