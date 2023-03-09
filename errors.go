package opsgenieclient

import "fmt"

type ValidationError struct {
	Err error
}

func (r *ValidationError) Error() string {
	return fmt.Sprintf("opsgenie: data validation error: %v", r.Err)
}
