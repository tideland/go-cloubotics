// Tideland Go Cloud Robotics - Options
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
// OPTION SETTER
//--------------------

// Option defines a function setting an option.
type Option func(*Cloud) error

// WithCloudProvider sets the cloud provider.
func WithCloudProvider(provider CloudProvider) Option {
	return func(c *Cloud) error {
		c.provider = provider
		return nil
	}
}

// EOF
