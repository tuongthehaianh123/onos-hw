// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package kpm

import (
	"context"
	"testing"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/stretchr/testify/assert"
	"github.com/tuongthehaianh123/onos-hw/pkg/manager"
	"github.com/tuongthehaianh123/onos-hw/test/utils"
)

// TestKpmSm is the function for Helmit-based integration test
func (s *TestSuite) TestKpmSm(t *testing.T) {
	cfg := manager.Config{
		CAPath:      "/tmp/tls.cacrt",
		KeyPath:     "/tmp/tls.key",
		CertPath:    "/tmp/tls.crt",
		ConfigPath:  "/tmp/config.json",
		E2tEndpoint: "onos-e2t:5150",
		GRPCPort:    5150,
		RicActionID: 10,
		SMName:      utils.KpmServiceModelName,
		SMVersion:   utils.KpmServiceModelVersion,
	}

	_, err := certs.HandleCertPaths(cfg.CAPath, cfg.KeyPath, cfg.CertPath, true)
	assert.NoError(t, err)

	mgr := manager.NewManager(cfg)
	mgr.Run()

	ctx, cancel := context.WithTimeout(context.Background(), utils.TestTimeout)
	defer cancel()

	err = utils.WaitForKPMIndicationMessages(ctx, t, mgr)
	assert.NoError(t, err)

	t.Log("KPM suite test passed")
}
