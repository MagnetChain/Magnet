// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	_ "embed"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//go:embed evm_debug_config.json
var EvmDebugConfig []byte

//go:embed evm_non_debug_config.json
var EvmNonDebugConfig []byte

var (
	// current avacloud settings
	LowGasLimit     = big.NewInt(12_000_000)
	MediumGasLimit  = big.NewInt(15_000_000) // C-Chain value
	HighGasLimit    = big.NewInt(20_000_000)
	LowTargetGas    = big.NewInt(25_000_000) // ~ 2.1x of gas limit
	MediumTargetGas = big.NewInt(45_000_000) // 3x of gas limit (also, 3x bigger than C-Chain)
	HighTargetGas   = big.NewInt(60_000_000) // 3x of gas limit

	NoDynamicFeesGasLimitToTargetGasFactor = big.NewInt(5)

	PrefundedEwoqAddress = common.HexToAddress("0x92028aA1c882641B27D5B32AB44Aeeb601AA92EC")
	PrefundedEwoqPrivate = "7963544318ec9f9483c2ae889d4aaa5af2a5f0e59dd52c7cd09d459f5bef1456"

	OneAvax                 = new(big.Int).SetUint64(1000000000000000000)
	defaultEVMAirdropAmount = new(big.Int).Exp(big.NewInt(21), big.NewInt(26), nil) // 21* 10^26
	defaultPoAOwnerBalance  = new(big.Int).Mul(OneAvax, big.NewInt(10))             // 10 Native Tokens
)
