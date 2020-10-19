package cabinet

import "math/big"

type PirateEthSetting struct {
	MBytesPerToken *big.Int
	PoolDeposit    *big.Int
	MinerDeposit   *big.Int
}

type PirateEthUserData struct {
	ChargeBalance *big.Int
	Epoch         *big.Int
}
