package cryptox

import (
	"crypto/rand"
	"math/big"
)

const (
	refLength = 10
	charset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateRef() string {
	b := make([]byte, refLength)
	max := big.NewInt(int64(len(charset)))
	for i := range b {
		for {
			num, err := rand.Int(rand.Reader, max)
			if err != nil {
				continue
			}
			b[i] = charset[num.Int64()]
			break
		}
	}
	return string(b)
}
