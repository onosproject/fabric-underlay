// SPDX-FileCopyrightText: 2023-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Package main is an entry point for launching the fabric underlay app
package main

import (
	"github.com/onosproject/fabric-underlay/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-net-lib/pkg/realm"
	"github.com/spf13/cobra"
)

var log = logging.GetLogger()

const (
	topoAddressFlag    = "topo-address"
	defaultTopoAddress = "onos-topo:5150"
)

// The main entry point
func main() {
	cmd := &cobra.Command{
		Use:  "fabric-underlay",
		RunE: runRootCommand,
	}
	realm.AddRealmFlags(cmd, "underlay")
	cmd.Flags().String(topoAddressFlag, defaultTopoAddress, "address:port or just :port of the onos-topo service")
	cli.AddServiceEndpointFlags(cmd, "fabric underlay gRPC")
	cli.Run(cmd)
}

func runRootCommand(cmd *cobra.Command, args []string) error {
	flags, err := cli.ExtractServiceEndpointFlags(cmd)
	if err != nil {
		return err
	}
	topoAddress, _ := cmd.Flags().GetString(topoAddressFlag)
	realmOptions := realm.ExtractOptions(cmd)

	log.Infow("Starting fabric-underlay", "realm-label", realmOptions.Label, "realm-value", realmOptions.Value)
	cfg := manager.Config{
		ServiceFlags: flags,
		TopoAddress:  topoAddress,
		RealmOptions: realmOptions,
	}
	return cli.RunDaemon(manager.NewManager(cfg))
}
