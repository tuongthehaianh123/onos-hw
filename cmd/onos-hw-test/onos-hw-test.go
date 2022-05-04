// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package main

import (
	"github.com/onosproject/helmit/pkg/registry"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/tuongthehaianh123/onos-hw/test/ha"
	"github.com/tuongthehaianh123/onos-hw/test/kpm"
)

func main() {
	registry.RegisterTestSuite("kpm", &kpm.TestSuite{})
	registry.RegisterTestSuite("ha", &ha.TestSuite{})
	test.Main()
}
