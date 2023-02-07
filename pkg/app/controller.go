// SPDX-FileCopyrightText: 2023-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Package app implements the fabric underlay control logic
package app

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

// Controller represents the link discovery control
type Controller struct {
}

// NewController creates a new fabric underlay app controller
func NewController() *Controller {
	ctrl := &Controller{}
	return ctrl
}

// Start starts the controller
func (c *Controller) Start() {
	log.Infof("Starting...")
	go c.run()
}

// Stop stops the controller
func (c *Controller) Stop() {
	log.Infof("Stopping...")
}

func (c *Controller) run() {
	log.Infof("Started")
	// TODO: implement me
	log.Infof("Stopped")
}
