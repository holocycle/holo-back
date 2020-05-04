package validator

import (
	go_validator "gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	Validator *go_validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validator: go_validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
