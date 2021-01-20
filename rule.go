package validation

//Rule ..
type Rule interface {
	Validate() (bool, error)
}

//Validate ..
func Validate(r *[]Rule) (bool, *[]error) {

	var errs []error
	valid := true

	for _, rule := range *r {
		ok, err := rule.Validate()

		if !ok {
			valid = false
			errs = append(errs, err)
		}
	}

	return valid, &errs
}
