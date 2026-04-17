package currency

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterCurrencyValidator(v *validator.Validate, t *ut.Translator) {
	_ = v.RegisterValidation("IsCorrectCurrency", isCorrectCurrency)
	_ = v.RegisterTranslation(
		"IsCorrectCurrency",
		*t,
		func(trans ut.Translator) error {
			if err := trans.Add("IsCorrectCurrency", "{0} is not correct currency", false); err != nil {
				return err
			}

			return nil
		},
		func(trans ut.Translator, fe validator.FieldError) string {
			msg, err := trans.T(fe.Tag(), fe.Field())
			if err != nil {
				return ""
			}

			return msg
		})
}

func isCorrectCurrency(fl validator.FieldLevel) bool {
	v := fl.Field().String()

	return Currency(v).IsValid()
}

func RegisterBlockchainNetworkValidator(v *validator.Validate, t *ut.Translator) {
	_ = v.RegisterValidation("IsCorrectBlockchainNetwork", isCorrectBlockchainNetwork)
	_ = v.RegisterTranslation(
		"IsCorrectBlockchainNetwork",
		*t,
		func(trans ut.Translator) error {
			if err := trans.Add("IsCorrectBlockchainNetwork", "{0} is not correct blockchainNetwork", false); err != nil {
				return err
			}

			return nil
		},
		func(trans ut.Translator, fe validator.FieldError) string {
			msg, err := trans.T(fe.Tag(), fe.Field())
			if err != nil {
				return ""
			}

			return msg
		})
}

func isCorrectBlockchainNetwork(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	return BlockchainNetwork(v).IsValid()
}
