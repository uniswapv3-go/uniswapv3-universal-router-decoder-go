package uniswapv3_universal_router_decoder_go

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/uniswapv3-go/uniswapv3-universal-router-decoder-go/command"
	"github.com/uniswapv3-go/uniswapv3-universal-router-decoder-go/uniswap"
	"math/big"
	"strings"
)

type Params struct {
	Commands []byte
	Inputs   [][]byte
	Deadline *big.Int
}

func Decode(data []byte) ([]command.Command, error) {
	urAbi, err := abi.JSON(strings.NewReader(uniswap.UniversalRouterMetaData.ABI))
	if err != nil {
		return nil, err
	}
	sigData := data[:4]

	method, err := urAbi.MethodById(sigData)
	if err != nil {
		return nil, err
	}
	argv, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return nil, err
	}
	if len(argv) != 3 {
		return nil, nil
	}
	var params Params
	err = method.Inputs.Copy(&params, argv)
	if err != nil {
		return nil, err
	}
	if len(params.Commands) != len(params.Inputs) {
		return nil, nil
	}
	var commands []command.Command
	for i := 0; i < len(params.Commands); i++ {
		commands = append(commands, command.Command{
			Command: params.Commands[i],
			Input:   params.Inputs[i],
		})
	}
	return commands, nil
}
