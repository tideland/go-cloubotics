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
	"context"

	"tideland.dev/go/cloubotics/types"
)

//--------------------
// CONSTANTS
//--------------------

const (
	// Version is the current version of the cloubotics package.
	Version = "0.1.0"
)

//--------------------
// CLOUD
//--------------------

// Cloud is the manager for cloud services.
type Cloud struct {
	ctx    context.Context
	cancel context.CancelFunc
	config *types.Config
}

// NewCloud creates a new cloud manager with the given options.
func NewCloud(ctx context.Context, config *types.Config) *Cloud {
	ctx, cancel := context.WithCancel(ctx)
	cloud := &Cloud{
		ctx:    ctx,
		cancel: cancel,
		config: config,
	}
	return cloud
}

// Stop stops the cloud manager.
func (c *Cloud) Stop() {
	c.cancel()
}

// Done allows to wait for the cloud manager to be stopped.
func (c *Cloud) Done() <-chan struct{} {
	return c.ctx.Done()
}

// Err returns an error if the cloud manager was stopped with an error.
func (c *Cloud) Err() error {
	return c.ctx.Err()
}

// Machines returns the selected machines of the cloud.
func (c *Cloud) Machines(selector types.Selector) types.Machiner {
	return c.config.Provider().Machines(selector)
}

// EOF
