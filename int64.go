package validation

import (
	"errors"
	"strconv"
)

//Int64Rule ..
type Int64Rule struct {
	Min   int64
	Max   int64
	Value string
}

//Validate ..
func (v *Int64Rule) Validate() (bool, error) {
	value, err := strconv.ParseInt(v.Value, 10, 16)

	if err != nil {
		return false, errors.New("Invalid")
	}

	if value < v.Min {
		return false, errors.New("Minimum is " + strconv.FormatInt(v.Min, 16))
	}

	if value > v.Max {
		return false, errors.New("Maximum is " + strconv.FormatInt(v.Max, 16))
	}

	return true, nil
}
