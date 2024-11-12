package random

import (
	"time"

	"math/rand"
)

func GenerateSixDigitsOTP() int {
	randomMechanic := rand.New(rand.NewSource(time.Now().UnixNano()))

    return 100000 + randomMechanic.Intn(900000)
}
