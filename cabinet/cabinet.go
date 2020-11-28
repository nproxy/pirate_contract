package cabinet

import "math/big"

type PirateEthSetting struct {
	MBytesPerToken *big.Int
	PoolDeposit    *big.Int
	MinerDeposit   *big.Int
}

func (pes *PirateEthSetting)Dup() *PirateEthSetting  {
	d:=&PirateEthSetting{
		MinerDeposit: &big.Int{},
		MBytesPerToken: &big.Int{},
		PoolDeposit: &big.Int{},
	}

	if pes.PoolDeposit == nil{
		pes.PoolDeposit = &big.Int{}
	}
	if pes.MinerDeposit == nil{
		pes.MinerDeposit = &big.Int{}
	}
	if pes.MBytesPerToken == nil{
		pes.MBytesPerToken = &big.Int{}
	}

	d.PoolDeposit.Add(d.PoolDeposit,pes.PoolDeposit)
	d.MinerDeposit.Add(d.MinerDeposit,pes.MinerDeposit)
	d.MBytesPerToken.Add(d.MBytesPerToken,pes.MBytesPerToken)

	return d
}

type PirateEthUserData struct {
	ChargeBalance *big.Int
	TotalTraffic  *big.Int
}
