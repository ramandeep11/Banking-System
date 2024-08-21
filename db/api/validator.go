package api

import (
	"simplebank/db/util"

	"github.com/go-playground/validator/v10"
)

// custom validator for the currency , so that we use it insted of oneof in json

var validCurrency validator.Func = func(feildLevel validator.FieldLevel) bool {

	if currency, ok := feildLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
