package v3

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
	Path             []byte
	Payer            common.Address
}

type SwapRouterExactOutputParams struct {
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	Path            []byte
	Payer           common.Address
}

func UnpackV3SwapExactIn(params *SwapRouterExactInputParams, input []byte) error {
	a, err := abi.JSON(strings.NewReader(SwapRouterMetaData.ABI))
	if err != nil {
		return err
	}
	method, ok := a.Methods["v3SwapExactInput"]
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

func UnpackV3SwapExactOut(params *SwapRouterExactOutputParams, input []byte) error {
	a, err := abi.JSON(strings.NewReader(SwapRouterMetaData.ABI))
	if err != nil {
		return err
	}
	method, ok := a.Methods["v3SwapExactOutput"]
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
