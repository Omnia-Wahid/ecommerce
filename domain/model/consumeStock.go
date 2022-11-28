package model

import "github.com/go-playground/validator"

type ConsumeStock struct {
	Country        string `json:"country" validate:"required"`
	ConsumedAmount *int   `json:"consumed_amount" validate:"required"`
}

func (cs *ConsumeStock) ValidateConsumeStock() error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(cs)
	if err != nil {
		//log.Println(err)
		return err
	}
	return nil
}
