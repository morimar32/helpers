package errors

type ValidationError struct {
	Err string
}

type DataAccessError struct {
	Err string
}

func (e *ValidationError) Error() string {
	return e.Err
}

func (e *DataAccessError) Error() string {
	return e.Err
}

func NewValidationError(err string) *ValidationError {
	ve := &ValidationError{
		Err: err,
	}
	return ve
}

func NewValidationErrorFromError(e error) *ValidationError {
	ve := &ValidationError{
		Err: e.Error(),
	}
	return ve
}

func NewDataAccessError(err string) *DataAccessError {
	dae := &DataAccessError{
		Err: err,
	}
	return dae
}
func NewDataAccessErrorFromError(e error) *DataAccessError {
	dae := &DataAccessError{
		Err: e.Error(),
	}
	return dae
}
