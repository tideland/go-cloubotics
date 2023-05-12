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
	"errors"

	"tideland.dev/go/cloubotics/services/machine"
)

//--------------------
// CONSTANTS
//--------------------

const (
	// Version is the current version of the cloubotics package.
	Version = "0.1.0"
)

var (
	// ErrNoCloudProvider is returned if the fresh created Cloud has
	// no CloudProvider set.
	ErrNoCloudProvider = errors.New("no cloud provider set")
)

//--------------------
// CLOUD PROVIDER
//--------------------

// CloudProvider is the interface for the individual cloud providers.
type CloudProvider interface {
	// ID returns the ID of the cloud provider implementation.
	ID() string
}

//--------------------
// CLOUD
//--------------------

// Cloud is the manager for cloud services.
type Cloud struct {
	provider CloudProvider
}

// NewCloud creates a new cloud manager with the given options.
func NewCloud(options ...Option) (*Cloud, error) {
	cloud := &Cloud{}
	for _, option := range options {
		if err := option(cloud); err != nil {
			return nil, err
		}
	}
	if cloud.provider == nil {
		return nil, ErrNoCloudProvider
	}
	return cloud, nil
}

// Machine returns the machine with the given ID.
func (c *Cloud) Machine(id machine.MachineID) (*machine.Machine, error) {
	return machine.NewMachine(id)
}

// EOF
