/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msptesttools

import (
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/mgmt"
	"github.com/hyperledger/fabric/orderer/common/localconfig"
)

// LoadTestMSPSetup sets up the local MSP
// and a chain MSP for the default chain
func LoadMSPSetupForTesting() error {
	localconf, err := localconfig.Load()
	if err != nil {
		panic(err)
	}
	localMSPDir := localconf.General.LocalMSPDir
	BCCSP := localconf.General.BCCSP
	localMSPID := localconf.General.LocalMSPID
	conf, err := msp.GetLocalMspConfig(localMSPDir, BCCSP, localMSPID)

	// dir := configtest.GetDevMspDir()
	// conf, err := msp.GetLocalMspConfig(dir, nil, "SampleOrg")

	if err != nil {
		return err
	}

	err = mgmt.GetLocalMSP(factory.GetDefault()).Setup(conf)
	if err != nil {
		return err
	}

	err = mgmt.GetManagerForChain("testchannelid").Setup([]msp.MSP{mgmt.GetLocalMSP(factory.GetDefault())})
	if err != nil {
		return err
	}

	return nil
}

// Loads the development local MSP for use in testing.  Not valid for production/runtime context
func LoadDevMsp() error {
	mspDir := configtest.GetDevMspDir()
	return mgmt.LoadLocalMsp(mspDir, nil, "SampleOrg")
}
