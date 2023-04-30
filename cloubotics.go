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

//--------------------
// TYPES
//--------------------

type VirtualMachineID = string

//--------------------
// MODEL INTERFACES
//--------------------

// VirtualMachine is the interface for a virtual machine.
type VirtualMachine interface {
	// ID returns the ID of the virtual machine.
	ID() string
}

//--------------------
// ACTIVITY INTERFACES
//--------------------

// VirtualMachineMonitor is the interface for the virtual machine monitor.
type VirtualMachineMonitor interface {
	// Set the creation callback for any new virtual machine.
	SetCreationCallback(callback func(vm VirtualMachine)) (VirtualMachineID, error)

	// Set the update callback for the given ID.
	SetUpdateCallback(id VirtualMachineID, callback func(vm VirtualMachine)) error

	// Set the deletion callback for the given ID.
	SetDeletionCallback(id VirtualMachineID, callback func(vm VirtualMachine)) error

	// Start starts the monitor.
	Start() error

	// Stop stops the monitor.
	Stop() error
}

// VirtualMachineService is the interface for the virtual machine service.
type VirtualMachineService interface {
	// VirtualMachineMonitor returns the virtual machine monitor.
	VirtualMachineMonitor() VirtualMachineMonitor

	// VirtualMachines returns the current virtual machines.
	VirtualMachines() ([]VirtualMachine, error)

	// VirtualMachine returns the current virtual machine with the given ID.
	VirtualMachine(id string) (VirtualMachine, error)

	// CreateVirtualMachine creates a new virtual machine.
	CreateVirtualMachine(vm VirtualMachine) (VirtualMachine, error)

	// UpdateVirtualMachine updates a virtual machine.
	UpdateVirtualMachine(vm VirtualMachine) (VirtualMachine, error)

	// DeleteVirtualMachine deletes a virtual machine.
	DeleteVirtualMachine(id string) error
}

// VolumeService is the interface for the volume service.
type VolumeService interface {
}

// BucketService is the interface for the bucket service.
type BucketService interface {
}

// DatabaseService is the interface for the database service.
type DatabaseService interface {
}

// Provider is the interface for the different cloud providers.
type Provider interface {
	// VirtualMachines returns the virtual machine service.
	VirtualMachines() VirtualMachineService

	// Volumes returns the volume service.
	Volumes() VolumeService

	// Buckets returns the bucket service.
	Buckets() BucketService

	// Databases returns the database service.
	Databases() DatabaseService
}

// EOF
