// Tideland Go Cloud Robotics - Types
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package types // import "tideland.dev/go/cloubotics/types"

//--------------------
// IMPORTS
//--------------------

import (
	"fmt"
)

//--------------------
// ERROR TYPES
//--------------------

// CloudError is the error type for the cloud package.
type CloudError struct {
	Msg string
	Err error
}

// NewCloudError is a convenient constructor for a CloudError.
func NewCloudError(err error, format string, args ...interface{}) *CloudError {
	return &CloudError{
		Msg: fmt.Sprintf(format, args...),
		Err: err,
	}
}

// Error implements the error interface.
func (e *CloudError) Error() string {
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s: %v", e.Msg, e.Err)
}

//--------------------
// MODEL TYPES
//--------------------

// ID is the type for IDs.
type ID = string

// Filter describes a filter for resources.
type Filter struct {
	// Type of the filter.
	Type string

	// Name of the filter.
	Name string

	// Values of the filter.
	Values []string
}

// Selector describes which resources should be selected.
type Selector struct {
	// IDs of the resources. An empty slice means all.
	IDs []ID

	// Filter for the resources. An empty filter means all.
	Filters []Filter
}

// Resource is common interface for all resources.
type Resource interface {
	// ID returns the ID of the resource.
	ID() ID
}

// Reconclier is the type for functions that reconcile something.
type Reconciler func(res Resource) error

// Machiner is the interface for the management of a set of machines.
type Machiner interface {
	// ID returns the ID of the machiner.
	ID() ID

	// Reconcile starts a reconciliation loop in the background.
	Reconcile(rec Reconciler) (Machiner, error)

	// Stop stops the reconciliation loop.
	Stop()

	// Err returns the error of a selection or inside the reconciliation
	// loop.
	Err() error
}

// EOF
