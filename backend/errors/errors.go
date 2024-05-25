package errors

import "errors"

func RequestInvalidError(errorMessage string) error {
	return errors.New(errorMessage)
}
