package magento2

import (
	"errors"
)

var ErrNoPointer = errors.New("target interface must be a pointer")

var ErrNotFound = errors.New("no document found")
