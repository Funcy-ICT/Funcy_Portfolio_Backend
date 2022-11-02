package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type ValidateErr struct {
	Tag   string `json:"tag"`
	Field string `json:"field"`
}

func Validate(s interface{}) ([]ValidateErr, error) {

	var e ValidateErr
	var errs []ValidateErr
	r := reflect.TypeOf(s)

	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			f, ok := r.FieldByName(err.Field())
			if !ok {
				return nil, errors.New("not found json field")
			}
			e = ValidateErr{
				Tag:   err.Tag(),
				Field: f.Tag.Get("json"),
			}
			errs = append(errs, e)
		}
		return errs, nil
	}
	return nil, nil
}
