package contract

import "errors"

// ContractErrors represents a list of ContractError.
type ContractErrors []ContractError

// Error returns a string representation of the ContractErrors to make sure it compatible with error type.
func (contractErrors *ContractErrors) Error() string {
	if contractErrors == nil {
		return ""
	}
	if len(*contractErrors) == 0 {
		return ""
	}
	var errs []error
	for _, ce := range *contractErrors {
		errs = append(errs, &ce)
	}
	err := errors.Join(errs...)
	if err == nil {
		return ""
	}
	return err.Error()
}

// ContractError represents an error with a message.
type ContractError struct {
	Message string
}

// Error returns the error message of the ContractError to make sure it compatible with error type.
func (e *ContractError) Error() string {
	return e.Message
}

var (
	// ErrContractCheckNotImplemented is an error indicating that the contract check is not implemented.
	ErrContractCheckNotImplemented ContractError = ContractError{"contract check not implemented"}
)
