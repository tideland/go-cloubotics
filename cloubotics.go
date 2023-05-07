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
	"tideland.dev/go/cloubotics/services/machine"
)

//--------------------
// CLOUD
//--------------------

// Cloud is the manager for cloud services.
type Cloud struct {
}

// NewCloud creates a new cloud manager.
func NewCloud() (*Cloud, error) {
	// TODO: Implement configuration handling
	return &Cloud{}, nil
}

// Machine returns the machine with the given ID.
func (c *Cloud) Machine(id machine.MachineID) (*machine.Machine, error) {
	return machine.NewMachine(id)
}

// EOF
