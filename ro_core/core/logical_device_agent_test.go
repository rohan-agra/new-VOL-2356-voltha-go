/*
 * Copyright 2019-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package core

import (
	"github.com/opencord/voltha-go/ro_core/config"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func MakeTestLogDevAgConfig() (*Core, error) {
	var core *Core
	var roCoreFlgs *config.ROCoreFlags
	var roC *roCore

	freePort, errP := freeport.GetFreePort()
	if errP == nil {
		freePortStr := strconv.Itoa(freePort)

		roCoreFlgs = config.NewROCoreFlags()
		roC = newROCore(roCoreFlgs)
		if (roC != nil) && (roCoreFlgs != nil) {
			addr := "127.0.0.1" + ":" + freePortStr
			cli, err := newKVClient("etcd", addr, 5)
			if err == nil {
				roC.kvClient = cli
				//return roCoreFlgs, roC
			}
		}
	}

	core = NewCore("ro_core", roCoreFlgs, roC.kvClient)

	return core, nil
}

// newLogicalDeviceAgent test
func TestNewLogicalDeviceAgent(t *testing.T) {
	core, _ := MakeTestLogDevAgConfig()
	assert.NotNil(t, core)
	logAgent := newLogicalDeviceAgent("log-dev", "", core.logicalDeviceMgr, core.deviceMgr, core.clusterDataProxy)
	assert.NotNil(t, logAgent)
}

// GetLogicalDevice, Flows, Ports test
func TestGetLogicalDevice(t *testing.T) {
	core, _ := MakeTestLogDevAgConfig()
	assert.NotNil(t, core)

	logDevMgr := newLogicalDeviceManager(core.deviceMgr, core.clusterDataProxy)
	assert.NotNil(t, logDevMgr)

	logAgent := newLogicalDeviceAgent("log-dev-ag", "", logDevMgr, core.deviceMgr, logDevMgr.clusterDataProxy)
	assert.NotNil(t, logAgent)

	logDevMgr.addLogicalDeviceAgentToMap(logAgent)

	logDev, err := logAgent.GetLogicalDevice()
	assert.Nil(t, logDev)
	assert.NotNil(t, err)

	Flws, err := logAgent.ListLogicalDeviceFlows()
	assert.Nil(t, Flws)
	assert.NotNil(t, err)
	FlwsGrp, err := logAgent.ListLogicalDeviceFlowGroups()
	assert.Nil(t, FlwsGrp)
	assert.NotNil(t, err)

	logDevPorts, err := logAgent.ListLogicalDevicePorts()
	assert.Nil(t, logDevPorts)
	assert.NotNil(t, err)

}
