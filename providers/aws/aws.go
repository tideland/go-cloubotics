// Tideland Go Cloud Robotics - Providers - AWS
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package aws // import "tideland.dev/go/cloubotics/providers/aws"

//--------------------
// IMPORTS
//--------------------

import (
	"context"
	"fmt"

	"tideland.dev/go/cloubotics/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

//--------------------
// PROVIDER
//--------------------

// Provider implements the Provider interface for AWS.
type Provider struct {
	ctx       context.Context
	config    *types.Config
	awsConfig aws.Config
}

// NewProvider creates a new AWS cloud provider.
func NewProvider(ctx context.Context) (*Provider, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}
	return &Provider{
		ctx:       ctx,
		awsConfig: cfg,
	}, nil
}

// ID returns the ID "aws" for this cloud provider implementation.
func (p *Provider) ID() types.ID {
	return "aws"
}

// SetConfig sets the configuration of the provider.
func (p *Provider) SetConfig(config *types.Config) {
	p.config = config
}

// Config returns the configuration of the provider.
func (p *Provider) Config() *types.Config {
	return p.config
}

// Machines returns the machines service.
func (p *Provider) Machines(selector types.Selector) types.Machiner {
	return newMachiner(p, selector)
}

// EOF
