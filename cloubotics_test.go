// Tideland Go Cloud Robotics
//
// Copyright (C) 2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package cloubotics_test

//--------------------
// IMPORTS
//--------------------

import (
	"context"
	"fmt"
	"tideland/go-cloubotics"
	"time"

	"tideland.dev/go/cloubotics/providers/aws"
	"tideland.dev/go/cloubotics/types"
)

//--------------------
// EXAMPLES
//--------------------

// Example shows the usage of the cloubotics package.
func Example() {
	ctx := context.Background()
	provider, err := aws.NewProvider(ctx)
	if err != nil {
		panic(err)
	}
	cloud, err := cloubotics.NewCloud(ctx, cloubotics.WithProvider(provider))
	if err != nil {
		panic(err)
	}
	svcs, err := cloud.Machines(types.Selector{
		IDs: []types.ID{},
		Filters: []types.Filter{
			{
				Type:   "tag",
				Name:   "env",
				Values: []string{"test"},
			},
		},
	}).Reconcile(func(res cloubotics.Resource) error {
		// Do something with the machine resource.
		fmt.Printf("machine: %v\n", res.ID())
		return nil
	})

	// Wait some time, then stop the reconciliation.

	time.Sleep(30 * time.Second)

	svcs.Stop()
}

// EOF
