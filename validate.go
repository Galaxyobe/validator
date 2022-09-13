package validator

import (
	"errors"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func StartWithAlpha(s string) error {
	if s == "" {
		return nil
	}
	r := []rune(s)
	if unicode.IsDigit(r[0]) {
		return errors.New("must start with alpha")
	}
	return nil
}

// IsStartWithAlpha Is start with alpha.
func IsStartWithAlpha(fl validator.FieldLevel) bool {
	return StartWithAlpha(fl.Field().String()) == nil
}
