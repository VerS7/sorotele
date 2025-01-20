package auth

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

// Генерация случайного лицевого счёта
func GenerateAccount(prefix string, length int) string {
	var code string = prefix

	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

// SHA256-хеширование с итерациями
func HashSHA256(data string, iterations int) [32]byte {
	hash := sha256.Sum256([]byte(data))
	for i := 0; i < iterations; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return hash
}

// SHA1-хеширование
func HashSHA1(data string) [20]byte {
	return sha1.Sum([]byte(data))
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
