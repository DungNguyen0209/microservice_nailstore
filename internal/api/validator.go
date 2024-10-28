package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/minhdung/nailstore/internal/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency
		util.IsSupportCurrency(currency)
	}
	return false
}
