// SPDX-FileCopyrightText: 2023-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Package manager contains the link agent manager coordinating lifecycle of the NB API and link discovery controller
package manager

import (
	"github.com/onosproject/fabric-underlay/pkg/app"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
)

var log = logging.GetLogger("manager")

// Config is a manager configuration
type Config struct {
	AgentUUID     string
	TargetAddress string
	ServiceFlags  *cli.ServiceEndpointFlags
}

// Manager is a single point of entry for the fabric-underlay
type Manager struct {
	cli.Daemon
	Config     Config
	controller *app.Controller
}

// NewManager initializes the application manager
func NewManager(cfg Config) *Manager {
	log.Infow("Creating manager")
	return &Manager{Config: cfg}
}

// Start initializes and starts the link controller and the NB gNMI API.
func (m *Manager) Start() error {
	log.Info("Starting Manager")

	// Initialize and start the app controller
	m.controller = app.NewController()
	m.controller.Start()

	// Starts NB server
	s := northbound.NewServer(cli.ServerConfigFromFlags(m.Config.ServiceFlags, northbound.SecurityConfig{}))
	s.AddService(logging.Service{})
	//s.AddService(nb.NewService(m.controller))
	return s.StartInBackground()
}

// Stop stops the manager
func (m *Manager) Stop() {
	log.Infow("Stopping Manager")
	m.controller.Stop()
}
