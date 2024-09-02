package main

import (
	"errors"
	"fmt"
)

// three fields:
// step to record the step name in an error;
// message msg that describes the
// condition; and a cause to store the underlying error that caused this step error:
var (
	ErrValidation = errors.New("Validation failed")
)

type stepErr struct {
	step  string
	msg   string
	cause error
}

func (s *stepErr) Error() string {
	return fmt.Sprint("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}

func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}
	return t.step == s.step
}
func (s *stepErr) Unwrap() error {
	return s.cause
}
