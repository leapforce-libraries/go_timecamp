package timecamp

import (
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
)

const (
	TimeFormat string = "15:04:05"
)

type TimeString civil.Time

func (d *TimeString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(TimeFormat, unquoted)
	if err != nil {
		return err
	}

	*d = TimeString(civil.TimeOf(_t))
	return nil
}

func (d *TimeString) ValuePtr() *civil.Time {
	if d == nil {
		return nil
	}

	_d := civil.Time(*d)
	return &_d
}

func (d TimeString) Value() civil.Time {
	return civil.Time(d)
}
