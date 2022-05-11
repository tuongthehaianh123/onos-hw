// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package utils

import (
	"context"
	"testing"

	hwAPI "github.com/onosproject/onos-api/go/onos/kpimon"
	"github.com/onosproject/onos-ric-sdk-go/pkg/utils/creds"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// HWServiceAddress defines the address and port for connections to the HW service
const HWServiceAddress = "onos-hw:5150"

// ConnectHWServiceHost connects to the onos HW service
func ConnectHWServiceHost() (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	return grpc.DialContext(context.Background(), HWServiceAddress, opts...)
}

// GetHWClient returns an SDK subscription client
func GetHWClient(t *testing.T) hwAPI.HWClient {
	conn, err := ConnectHWServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return hwAPI.NewHWClient(conn)
}
