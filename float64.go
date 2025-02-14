package validation

import (
	"errors"
	"strconv"
)

//Float64Rule ..
type Float64Rule struct {
	Name  string
	Min   float64
	Max   float64
	Value string
}

//Validate ..
func (v *Float64Rule) Validate() (bool, error) {
	value, err := strconv.ParseFloat(v.Value, 64)

	if err != nil {
		return false, errors.New(v.Name + " Invalid")
	}

	if value < v.Min {
		return false, errors.New(v.Name + " Minimum is " + strconv.FormatFloat(v.Min, 'f', -1, 64))
	}

	if value > v.Max {
		return false, errors.New(v.Name + " Maximum is " + strconv.FormatFloat(v.Max, 'f', -1, 64))
	}

	return true, nil
}
