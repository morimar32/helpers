package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func TranslateErrorTogRPCStatusError(err error) error {
	if err != nil {
		if errors.Is(err, &DataAccessError{}) {
			return status.Errorf(codes.Internal, err.Error())
		}
		if errors.Is(err, &ValidationError{}) {
			return status.Errorf(codes.InvalidArgument, err.Error())
		}
		return status.Errorf(codes.Unknown, err.Error())
	}
	return nil
}
