// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "strings"

// UnimplementedError is used to indicate that the operation
// being performed is not yet implemented. It is primarily used
// to permit execution of code generation to continue whilst
// logging any unimplemented operations.
type UnimplementedError struct {
	err  error
	path []string
}

// Error returns the underlying error string.
func (e *UnimplementedError) Error() string {
	return e.err.Error()
}

// Path returns a dot-separated path.
func (e *UnimplementedError) Path() string {
	return strings.Join(e.path, ".")
}

// NewUnimplementedError returns an UnimplementedError populated with the
// supplied error and path.
func NewUnimplementedError(err error, path ...string) *UnimplementedError {
	return &UnimplementedError{
		err:  err,
		path: path,
	}
}

// NestedUnimplementedError returns an UnimplementedError with a path that includes
// the supplied parentPath.
func (e *UnimplementedError) NestedUnimplementedError(parentPath string) *UnimplementedError {
	newErr := &UnimplementedError{
		err:  e.err,
		path: append([]string{parentPath}, e.path...),
	}

	return newErr
}
