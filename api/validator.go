package api

import (
	"github.com/Elessurgon/go_pract/util"
	"github.com/go-playground/validator/v10"
)


var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
    if currency, ok := fieldLevel.Field().Interface().(string); ok {
        return util.IsSupportedCurrency(currency)
    }
    return false
}