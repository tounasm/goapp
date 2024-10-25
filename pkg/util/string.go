package util

import (
	"math/rand"
)

var randx = rand.NewSource(42)

// RandHexString returns a random hex string of length n.
func RandHexString(n int) string {
	const letterBytes = "0123456789ABCDEF"
	const (
		letterIdxBits = 4                    // 4 bits to represent a hex digit index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of hex digit indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax hex digits!
	for i, cache, remain := n-1, randx.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randx.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
