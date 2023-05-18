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

// Provider is the interface for the individual cloud providers.
type Provider interface {
	// ID returns the ID of the cloud provider implementation.
	ID() types.ID

	// Machines returns the machiner for the selected machines.
	Machines(selector types.Selector) types.Machiner
}

// EOF
