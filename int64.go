package validation

import (
	"errors"
	"strconv"
)

//Int64Rule ..
type Int64Rule struct {
	Name  string
	Min   int64
	Max   int64
	Value string
}

//Validate ..
func (v *Int64Rule) Validate() (bool, error) {
	value, err := strconv.ParseInt(v.Value, 10, 16)

	if err != nil {
		return false, errors.New(v.Name + " Invalid")
	}

	if value < v.Min {
		return false, errors.New(v.Name + " Minimum is " + strconv.FormatInt(v.Min, 10))
	}

	if value > v.Max {
		return false, errors.New(v.Name + " Maximum is " + strconv.FormatInt(v.Max, 10))
	}

	return true, nil
}
