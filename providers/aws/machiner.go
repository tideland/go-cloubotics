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
	"time"

	"tideland.dev/go/cloubotics/types"
	"tideland.dev/go/uuid"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

//--------------------
// CONSTANTS
//--------------------

const (
	// machinerInterval is the interval for the machiner.
	machinerInterval = 30 * time.Second
)

//--------------------
// MACHINE
//--------------------

// machine implements the Resource interface for AWS.
type machine struct {
	provider *Provider
	id       types.ID
	instance ec2types.Instance
}

// ID implements the Resource interface.
func (m *machine) ID() types.ID {
	return m.id
}

// State implements the Resource interface.
func (m *machine) State() types.ResourceState {
	switch m.instance.State.Name {
	case ec2types.InstanceStateNamePending:
		return types.ResourcePending
	case ec2types.InstanceStateNameRunning:
		return types.ResourceRunning
	case ec2types.InstanceStateNameStopping:
		return types.ResourceStopping
	case ec2types.InstanceStateNameStopped:
		return types.ResourceStopped
	case ec2types.InstanceStateNameShuttingDown:
		return types.ResourceShuttingDown
	case ec2types.InstanceStateNameTerminated:
		return types.ResourceTerminated
	}
	return types.ResourceUnknown
}

//--------------------
// MACHINER
//--------------------

// machiner implements the machiner interface for AWS.
type machiner struct {
	provider *Provider
	ctx      context.Context
	cancel   context.CancelFunc
	client   *ec2.Client
	id       types.ID
	input    *ec2.DescribeInstancesInput
}

// newMachiner creates a new AWS machiner and prepares the selection
// of the machines as goroutine.
func newMachiner(provider *Provider, selector types.Selector) *machiner {
	ctx, cancel := context.WithCancel(provider.ctx)
	input := &ec2.DescribeInstancesInput{}
	if len(selector.IDs) > 0 {
		input.InstanceIds = selector.IDs
	}
	if len(selector.Filters) > 0 {
		input.Filters = make([]ec2types.Filter, len(selector.Filters))
		for i, f := range selector.Filters {
			input.Filters[i] = ec2types.Filter{
				Name:   aws.String(f.Type + ":" + f.Name),
				Values: f.Values,
			}
		}
	}
	return &machiner{
		provider: provider,
		ctx:      ctx,
		cancel:   cancel,
		client:   ec2.NewFromConfig(provider.awsConfig),
		id:       uuid.New().String(),
		input:    input,
	}
}

// ID implements the Machiner interface.
func (m *machiner) ID() types.ID {
	return m.id
}

// Reconile implements the Machiner interface. It starts a goroutine
// that reconciles the machines.
func (m *machiner) Reconcile(rec types.Reconciler) (types.Machiner, error) {
	go m.reconcile(rec)
	return m, nil
}

// reconcile is the goroutine that reconciles the machines.
func (m *machiner) reconcile(rec types.Reconciler) {
	ticker := time.NewTicker(machinerInterval)
	defer ticker.Stop()
	for {
		select {
		case <-m.ctx.Done():
			m.provider.config.Logger().Printf("machiner %q stopped", m.id)
			return
		case <-ticker.C:
			err := m.process(rec)
			if err != nil {
				m.provider.config.Logger().Printf("machiner %q failed: %v", m.id, err)
				return
			}
		}
	}
}

// process processes the machines.
func (m *machiner) process(rec types.Reconciler) error {
	m.provider.config.Logger().Printf("machiner %q processing", m.id)
	output, err := m.client.DescribeInstances(m.ctx, m.input)
	if err != nil {
		return err
	}
	for _, res := range output.Reservations {
		for _, inst := range res.Instances {
			err := rec(&machine{
				provider: m.provider,
				id:       types.ID(*inst.InstanceId),
				instance: inst,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Stop implements the Machiner interface.
func (m *machiner) Stop() {
	m.cancel()
}

// Err implements the Machiner interface.
func (m *machiner) Err() error {
	return m.ctx.Err()
}

// EOF
