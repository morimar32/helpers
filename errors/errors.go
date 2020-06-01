package errors

type ValidationError struct {
	err string
}

type DataAccessError struct {
	err string
}

func (e *ValidationError) Error() string {
	return e.err
}

func (e *DataAccessError) Error() string {
	return e.err
}

func NewValidationError(e error) *ValidationError {
	ve := &ValidationError{
		err: e.Error(),
	}
	return ve
}

func NewDataAccessError(e error) *DataAccessError {
	dae := &DataAccessError{
		err: e.Error(),
	}
	return dae
}
