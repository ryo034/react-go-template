package authentication

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

const maxDigits = 6

func GenerateOTPCode() (string, error) {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0*d", maxDigits, bi), nil
}
