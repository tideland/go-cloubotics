// Tideland Go Cloud Robotics - Services - Machine
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package machine // import "tideland.dev/go/cloubotics/services/machine"

//--------------------
// IMPORTS
//--------------------

//--------------------
// TYPES
//--------------------

// MachineID is the ID of a machine.
type MachineID = string

//--------------------
// MACHINE
//--------------------

// Machine is the interface for a machine.
type Machine struct {
	id MachineID
}

// NewMachine creates a new machine representation.
func NewMachine(id MachineID) (*Machine, error) {
	// TODO: Implement
	return &Machine{
		id: id,
	}, nil
}

// ID returns the ID of the machine.
func (m *Machine) ID() MachineID {
	return m.id
}

// EOF
