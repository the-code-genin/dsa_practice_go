package algorithms

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

func GenerateRandomNumber(max int64) (int64, error) {
	data := make([]byte, 8)
	_, err := rand.Read(data)
	if err != nil {
		return 0, err
	}

	val, err := rand.Int(bytes.NewReader(data), big.NewInt(max))
	if err != nil {
		return 0, err
	}

	return val.Int64(), nil
}
