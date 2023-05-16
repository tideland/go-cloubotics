// Tideland Go Cloud Robotics
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package cloubotics // import "tideland.dev/go/cloubotics"

//--------------------
// IMPORTS
//--------------------

import (
	"tideland.dev/go/cloubotics/types"
)

//--------------------
// INTERFACES
//--------------------

// Resource is common interface for all resources.
type Resource interface {
	// ID returns the ID of the resource.
	ID() types.ID
}

// Reconclier is the type for functions that reconcile something.
type Reconciler func(res Resource) error

// Services defines the common interface for all services.
type Services interface {
	// Reconcile starts a reconciliation loop in the background.
	Reconcile(rec Reconciler) (Services, error)

	// Stop stops the reconciliation loop.
	Stop()

	// Err returns the error of a selection or inside the reconciliation
	// loop.
	Err() error
}

// Machines is the interface for the machines service.
type Machines interface {
	Services
}

// EOF
