package geomgen

import (
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
)

func SeedFromString(s string) int64 {
	h := sha256.Sum256([]byte(s))
	return int64(binary.BigEndian.Uint64(h[:8]))
}

func NewRand(seedStr string) *rand.Rand {
	return rand.New(rand.NewSource(SeedFromString(seedStr)))
}
