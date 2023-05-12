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

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

//--------------------
// CLOUD PROVIDER
//--------------------

// CloudProvider implements the CloudProvider interface for AWS.
type CloudProvider struct {
	cfg config.Config
}

// NewCloudProvider creates a new AWS cloud provider.
func NewCloudProvider(ctx context.Context) (*CloudProvider, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}
	return &CloudProvider{
		cfg: cfg,
	}, nil
}

// ID returns the ID "aws" for this cloud provider implementation.
func (p *CloudProvider) ID() string {
	return "aws"
}

// EOF
