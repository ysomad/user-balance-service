package apperror

import (
	"fmt"
)

const format = "%s: %w"

func New(msg string, err error) error {
	return fmt.Errorf(format, msg, err)
}
