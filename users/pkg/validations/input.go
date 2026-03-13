package validations

import (
	"errors"
	"net/mail"
	"unicode"
)

func IsValidEmailSyntax(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	if !hasUpper {
		return errors.New("la contraseña debe tener al menos una letra mayúscula")
	}
	if !hasSpecial {
		return errors.New("la contraseña debe tener al menos un carácter especial")
	}

	return nil
}
