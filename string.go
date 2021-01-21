package validation

import (
	"errors"
	"strconv"
)

//StringRule ..
type StringRule struct {
	Name  string
	Min   int
	Max   int
	Value string
}

//Validate ..
func (v *StringRule) Validate() (bool, error) {
	if len(v.Value) < v.Min {
		return false, errors.New(v.Name + " Minimum Length is " + strconv.FormatInt(int64(v.Min), 10))
	}

	if len(v.Value) > v.Max {
		return false, errors.New(v.Name + " Maximum Length is " + strconv.FormatInt(int64(v.Max), 10))
	}

	return true, nil
}
