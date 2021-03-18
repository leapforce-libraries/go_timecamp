package timecamp

import (
	"fmt"
	"strconv"
	"strings"
)

type BoolString bool

func (bl *BoolString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		*bl = false
		return nil
	}

	i, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}

	if i == 0 {
		*bl = false
		return nil
	}

	if i == 1 {
		*bl = true
		return nil
	}

	return fmt.Errorf("Cannot parse '%s' to BoolString", string(b))
}

func (bl *BoolString) ValuePtr() *bool {
	if bl == nil {
		return nil
	}

	_b := bool(*bl)
	return &_b
}

func (bl BoolString) Value() bool {
	return bool(bl)
}
