// Tideland Go Cloud Robotics - Types
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package types // import "tideland.dev/go/cloubotics/types"

//--------------------
// CLOUD
//--------------------

// ResourceState is the type for the state of a resource.
type ResourceState string

const (
	// The different states of a resource.
	ResourceUnknown      ResourceState = "unknown"
	ResourcePending      ResourceState = "pending"
	ResourceRunning      ResourceState = "running"
	ResourceStopping     ResourceState = "stopping"
	ResourceStopped      ResourceState = "stopped"
	ResourceShuttingDown ResourceState = "shutting-down"
	ResourceTerminated   ResourceState = "terminated"
)

// Resource is common interface for all resources.
type Resource interface {
	// ID returns the ID of the resource.
	ID() ID

	// State returns the state of the resource.
	State() ResourceState

	// Internal returns an internal value defined by kind and name.
	// The caller must know the type of the value to assert it.
	Internal(kind, name string) (any, error)
}

// Reconclier is the type for functions that reconcile something.
type Reconciler func(res Resource) error

// Provider is the interface for the individual cloud providers.
type Provider interface {
	// ID returns the ID of the cloud provider implementation.
	ID() ID

	// SetConfig allows to set the configuration.
	SetConfig(cfg *Config)

	// Config returns the configuration.
	Config() *Config

	// Machines returns the machiner for the selected machines.
	Machines(selector Selector) Machiner
}

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
