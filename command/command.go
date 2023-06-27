package command

const (
	V3_SWAP_EXACT_IN = iota
	V3_SWAP_EXACT_OUT
	PERMIT2_TRANSFER_FROM
	PERMIT2_PERMIT_BATCH
	SWEEP
	TRANSFER
	PAY_PORTION
	_
	V2_SWAP_EXACT_IN
	V2_SWAP_EXACT_OUT
	PERMIT2_PERMIT
	WRAP_ETH
	UNWRAP_WETH
	PERMIT2_TRANSFER_FROM_BATCH
	_
	_
	SEAPORT_V1_5
	LOOKS_RARE_721
	NFTX
	CRYPTOPUNKS
	LOOKS_RARE_1155
	OWNER_CHECK_721
	OWNER_CHECK_1155
	SWEEP_ERC721
	X2Y2_721
	SUDOSWAP
	NFT20
	X2Y2_1155
	FOUNDATION
	SWEEP_ERC1155

	MASK = 0x3F
)

func Revertable(command byte) bool {
	return 0x1&command == 0x1
}

func RealCommand(command byte) byte {
	return command & MASK
}

func RealCommandName(command byte) string {
	switch command & MASK {
	case V3_SWAP_EXACT_IN:
		return "V3_SWAP_EXACT_IN"
	case V3_SWAP_EXACT_OUT:
		return "V3_SWAP_EXACT_OUT"
	case PERMIT2_TRANSFER_FROM:
		return "PERMIT2_TRANSFER_FROM"
	case PERMIT2_PERMIT_BATCH:
		return "PERMIT2_PERMIT_BATCH"
	case SWEEP:
		return "SWEEP"
	case TRANSFER:
		return "TRANSFER"
	case PAY_PORTION:
		return "PAY_PORTION"
	case V2_SWAP_EXACT_IN:
		return "V2_SWAP_EXACT_IN"
	case V2_SWAP_EXACT_OUT:
		return "V2_SWAP_EXACT_OUT"
	case PERMIT2_PERMIT:
		return "WRAP_ETH"
	case UNWRAP_WETH:
		return "UNWRAP_WETH"
	case PERMIT2_TRANSFER_FROM_BATCH:
		return "PERMIT2_TRANSFER_FROM_BATCH"
	case SEAPORT_V1_5:
		return "SEAPORT_V1_5"
	case LOOKS_RARE_721:
		return "LOOKS_RARE_721"
	case NFTX:
		return "NFTX"
	case CRYPTOPUNKS:
		return "CRYPTOPUNKS"
	case LOOKS_RARE_1155:
		return "LOOKS_RARE_1155"
	case OWNER_CHECK_721:
		return "OWNER_CHECK_721"
	case OWNER_CHECK_1155:
		return "OWNER_CHECK_1155"
	case SWEEP_ERC721:
		return "SWEEP_ERC721"
	case X2Y2_721:
		return "X2Y2_721"
	case SUDOSWAP:
		return "SUDOSWAP"
	case NFT20:
		return "NFT20"
	case X2Y2_1155:
		return "X2Y2_1155"
	case FOUNDATION:
		return "FOUNDATION"
	case SWEEP_ERC1155:
		return "SWEEP_ERC1155"

	default:
		return ""
	}
}

type Command struct {
	Command byte
	Input   []byte
}
