package account_helper

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func PrivateKeyFromStr(keyStr string) ecdsa.PrivateKey{
	var pri ecdsa.PrivateKey
	pri.D, _ = new(big.Int).SetString(keyStr, 16)
	pri.PublicKey.Curve = crypto.S256()
	pri.PublicKey.X, pri.PublicKey.Y = pri.PublicKey.Curve.ScalarBaseMult(pri.D.Bytes())
	return pri
}
