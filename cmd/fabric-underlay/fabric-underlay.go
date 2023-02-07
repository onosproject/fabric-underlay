// SPDX-FileCopyrightText: 2023-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Package main is an entry point for launching the fabric underlay app
package main

import (
	"github.com/onosproject/fabric-underlay/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/spf13/cobra"
)

var log = logging.GetLogger()

// The main entry point
func main() {
	cmd := &cobra.Command{
		Use:  "fabric-underlay",
		RunE: runRootCommand,
	}
	cli.AddServiceEndpointFlags(cmd, "fabric underlay gRPC")
	cli.Run(cmd)
}

func runRootCommand(cmd *cobra.Command, args []string) error {
	flags, err := cli.ExtractServiceEndpointFlags(cmd)
	if err != nil {
		return err
	}

	log.Infof("Starting fabric-underlay")
	cfg := manager.Config{
		ServiceFlags: flags,
	}
	return cli.RunDaemon(manager.NewManager(cfg))
}
