package zh

import (
	"testing"

	. "github.com/go-playground/assert/v2"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	myvalidator "github.com/galaxyobe/validator"
)

func TestTranslations(t *testing.T) {
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	myvalidator.RegisterValidators(validate)

	err := RegisterTranslations(validate, trans)
	Equal(t, err, nil)

	type Test struct {
		StartWithAlphaString string `validate:"startwithalpha"`
	}

	var test = Test{
		StartWithAlphaString: "0123abc",
	}

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(validator.ValidationErrors)
	Equal(t, ok, true)

	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "Test.StartWithAlphaString",
			expected: "StartWithAlphaString必须以字母开头",
		},
	}

	for _, tt := range tests {
		var fe validator.FieldError
		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fe = e
				break
			}
		}
		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}
}
