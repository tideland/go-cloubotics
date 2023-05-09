// Tideland Go Cloud Robotics - Providers - AWS
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
// CLOUD PROVIDER
//--------------------

// CloudProvider implements the CloudProvider interface for AWS.
type CloudProvider struct{}

// NewCloudProvider creates a new AWS cloud provider.
func NewCloudProvider() (*CloudProvider, error) {
	return &CloudProvider{}, nil
}

// ID returns the ID "aws" for this cloud provider implementation.
func (p *CloudProvider) ID() string {
	return "aws"
}

// EOF
