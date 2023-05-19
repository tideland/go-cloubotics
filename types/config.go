// Tideland Go Cloud Robotics - Types
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package types // import "tideland.dev/go/cloubotics/types"

//--------------------
// IMPORTS
//--------------------

import (
	"log"
)

//--------------------
// CONFIGURATION
//--------------------

// Config contains the global configuration.
type Config struct {
	// Provider is the cloud provider to use.
	provider Provider

	// Logger is the logger to use.
	logger Logger
}

// NewConfig creates a new configuration with the given options.
func NewConfig(options ...Option) (*Config, error) {
	cfg := &Config{}
	for _, option := range options {
		if err := option(cfg); err != nil {
			return nil, err
		}
	}
	// Check the configured fields.
	if cfg.provider == nil {
		return nil, NewCloudError(nil, "no cloud provider defined")
	}
	if cfg.logger == nil {
		cfg.logger = log.Default()
	}
	return cfg, nil
}

// Provider returns the configured cloud provider.
func (c *Config) Provider() Provider {
	return c.provider
}

// Logger returns the configured logger.
func (c *Config) Logger() Logger {
	return c.logger
}

// Configurator is a function returning a configuration. It is used
// to pass the top level configuration to all using types.
type Configurator func() *Config

//--------------------
// OPTION SETTER
//--------------------

// Option defines a function setting an option.
type Option func(*Config) error

// WithProvider sets the cloud provider.
func WithProvider(provider Provider) Option {
	return func(c *Config) error {
		c.provider = provider
		return nil
	}
}

// WithLogger sets the logger.
func WithLogger(logger Logger) Option {
	return func(c *Config) error {
		c.logger = logger
		return nil
	}
}

// EOF
