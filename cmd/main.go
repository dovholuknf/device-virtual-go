// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/edgexfoundry/device-sdk-go/v3/pkg/startup"

	"github.com/edgexfoundry/device-virtual-go"
	"github.com/edgexfoundry/device-virtual-go/internal/driver"
)

const (
	serviceName string = "device-virtual"
)

func main() {
	fmt.Println("AAAAA")
	fmt.Println("AAAAA")
	fmt.Println("AAAAA")
	fmt.Println("bbbbb")
	fmt.Println("bbbbb")
	d := driver.NewVirtualDeviceDriver()
	startup.Bootstrap(serviceName, device_virtual.Version, d)
}
