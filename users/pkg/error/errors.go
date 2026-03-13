package apperror

import "errors"

var (
	ErrNoChanges    = errors.New("no se detectaron cambios en el recurso")
	ErrEmptyInput   = errors.New("el campo esta vacio, ingresa el campo: ")
	ErrInvalidEmail = errors.New("email invalido")
)
