package model

// Validator for customer
func (a *Product) Validate() error {
	errs := ValidationErrors{}
	return errs.AsError()
}
