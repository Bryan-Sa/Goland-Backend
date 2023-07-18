package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Erreur de validation pour le champ %s: %s\n", err.Field(), err.Tag())
		}
		return errors.New("données de requête invalides")
	}
	return nil
}
