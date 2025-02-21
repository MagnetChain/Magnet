// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subnet

import (
	"fmt"

	"github.com/ava-labs/avalanche-cli/tests/e2e/commands"
	"github.com/ava-labs/avalanche-cli/tests/e2e/utils"
	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

const (
	CLIBinary         = "./bin/avalanche"
	keyName           = "ewoq"
	ewoqEVMAddress    = "0x92028aA1c882641B27D5B32AB44Aeeb601AA92EC"
	ewoqPChainAddress = "P-custom18jma8ppw3nhx5r4ap8clazz0dps7rv5u9xde7p"
)

var _ = ginkgo.Describe("[Etna Add Validator SOV Local]", func() {
	ginkgo.It("Create Etna Subnet Config", func() {
		commands.CreateEtnaSubnetEvmConfig(
			utils.SubnetName,
			ewoqEVMAddress,
			commands.PoS,
		)
	})
	ginkgo.It("Can deploy blockchain to localhost and upsize it", func() {
		output := commands.StartNetwork()
		fmt.Println(output)
		output, err := commands.DeployEtnaBlockchain(
			utils.SubnetName,
			"",
			nil,
			ewoqPChainAddress,
			false, // convertOnly
		)
		gomega.Expect(err).Should(gomega.BeNil())
		fmt.Println(output)
		output, err = commands.AddEtnaSubnetValidatorToCluster(
			"",
			utils.SubnetName,
			"",
			ewoqPChainAddress,
			1,
			1,
			true,
		)
		gomega.Expect(err).Should(gomega.BeNil())
		fmt.Println(output)
	})

	ginkgo.It("Can destroy local node", func() {
		output, err := commands.DestroyLocalNode(utils.TestLocalNodeName)
		gomega.Expect(err).Should(gomega.BeNil())
		fmt.Println(output)
	})

	ginkgo.It("Can destroy Etna Local Network", func() {
		commands.CleanNetwork()
	})

	ginkgo.It("Can remove Etna Subnet Config", func() {
		commands.DeleteSubnetConfig(utils.SubnetName)
	})
})
