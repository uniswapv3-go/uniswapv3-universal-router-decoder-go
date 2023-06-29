package v2

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type SwapRouterExactInputParams struct {
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	Path             []common.Address
	Payer            common.Address
}

type SwapRouterExactOutputParams struct {
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	Path            []common.Address
	Payer           common.Address
}

func UnpackV2SwapExactIn(params *SwapRouterExactInputParams, input []byte) error {
	a, err := abi.JSON(strings.NewReader(SwapRouterMetaData.ABI))
	if err != nil {
		return err
	}
	method, ok := a.Methods["v2SwapExactInput"]
	if !ok {
		return errors.New("no this method")
	}
	argv, err := method.Inputs.Unpack(input)
	if err != nil {
		return err
	}
	err = method.Inputs.Copy(params, argv)
	if err != nil {
		return err
	}
	return nil
}

func UnpackV2SwapExactOut(params *SwapRouterExactOutputParams, input []byte) error {
	a, err := abi.JSON(strings.NewReader(SwapRouterMetaData.ABI))
	if err != nil {
		return err
	}
	method, ok := a.Methods["v2SwapExactOutput"]
	if !ok {
		return errors.New("no this method")
	}
	argv, err := method.Inputs.Unpack(input)
	if err != nil {
		return err
	}
	err = method.Inputs.Copy(params, argv)
	if err != nil {
		return err
	}
	return nil
}
