package api

import (
	"fmt"

	"github.com/ThoPham02/simp_bank/util"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorError(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			return fmt.Sprintf("%s is required", err.Field())
		case "email":
			return fmt.Sprintf("%s is not a valid email address", err.Field())
		case "min":
			return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
		case "max":
			return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
		case "currency":
			return fmt.Sprintf("%s invalid", err.Field())
		case "gte":
			return fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
		case "lte":
			return fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
		}
	}
	return err.Error()
}

var ValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsValidCurrency(currency)
	}
	return false
}
