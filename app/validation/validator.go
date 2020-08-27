package validation

import (
	"errors"
	"github.com/go-playground/validator"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Check(val interface{}) (err error) {
	err = validate.Struct(val)
	if err != nil {
		var errSlices []string
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				v := FieldError{e}.String()
				errSlices = append(errSlices, v)
			}
		}
		err = errors.New(strings.Join(errSlices, ";"))
		return err
	}
	return
}
