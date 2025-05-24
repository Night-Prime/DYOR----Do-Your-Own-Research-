package errors

import "fmt"

// my solution to solving the Error Propagation Problem:

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// Using this not for only DB errors but package ones too
type DatabaseError struct{
	Message string
	Err		error
}

func (e * DatabaseError) Error() string {
	if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

// Using this for custom errors too (buggy until proven otherwise)
type CustomError struct {
    Message string
    Value   string
    Err     error
}

func (e *CustomError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s (value: %s): %v", e.Message, e.Value, e.Err)
    }
    if e.Value != "" {
        return fmt.Sprintf("%s (value: %s)", e.Message, e.Value)
    }
    return e.Message
}
