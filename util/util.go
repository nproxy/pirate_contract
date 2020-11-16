package util

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func BalanceHuman(balance *big.Int) float64 {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	v := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	vv, _ := v.Float64()
	return vv
}

func BalanceEth(balance float64) *big.Int {
	fbalance := new(big.Float)
	fbalance.SetFloat64(balance)
	v := new(big.Float).Mul(fbalance, big.NewFloat(math.Pow10(18)))

	vv := new(big.Int)
	v.Int(vv)

	return vv
}

func TrafficGBytes(traffic *big.Int) float64 {
	f := new(big.Float)
	f.SetString(traffic.String())
	v := new(big.Float).Quo(f, big.NewFloat(math.Pow10(9)))

	vv, _ := v.Float64()

	return vv
}

func TrafficMBytes(traffic *big.Int) float64 {
	f := new(big.Float)
	f.SetString(traffic.String())
	v := new(big.Float).Quo(f, big.NewFloat(math.Pow10(6)))

	vv, _ := v.Float64()

	return vv
}

func Float2String(f float64, point int) string {
	return fmt.Sprintf("%."+strconv.Itoa(point)+"f", f)
}

func MaxBigInt(x, y *big.Int) *big.Int {
	if x == nil && y == nil {
		return &big.Int{}
	}

	if x == nil {
		return y
	}
	if y == nil {
		return x
	}

	if x.Cmp(y) > 0 {
		return x
	}

	return y
}
