package timecamp

import (
	"encoding/json"
	"fmt"
)

type BoolInt bool

func (bl *BoolInt) UnmarshalJSON(b []byte) error {
	var i int

	err := json.Unmarshal(b, &i)
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

	return fmt.Errorf("Cannot parse '%s' to BoolInt", string(b))
}

func (bl *BoolInt) ValuePtr() *bool {
	if bl == nil {
		return nil
	}

	_b := bool(*bl)
	return &_b
}

func (bl BoolInt) Value() bool {
	return bool(bl)
}
