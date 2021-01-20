package validation

import (
	"errors"
	"strconv"
)

//StringRule ..
type StringRule struct {
	Min   int
	Max   int
	Value string
}

//Validate ..
func (v *StringRule) Validate() (bool, error) {
	if len(v.Value) < v.Min {
		return false, errors.New("Minimum Length is " + strconv.FormatInt(int64(v.Min), 16))
	}

	if len(v.Value) > v.Max {
		return false, errors.New("Maximum Length is " + strconv.FormatInt(int64(v.Max), 16))
	}

	return true, nil
}
