package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

// SHA256-хеширование с итерациями
func HashSHA256(data string, iterations int) [32]byte {
	hash := sha256.Sum256([]byte(data))
	for i := 0; i < iterations; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return hash
}

// Сравнение хэшей
func CompareHashes(first, second []byte) bool {
	if compare := bytes.Compare(first, second); compare == 0 {
		return true
	}
	return false
}

func HashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}
