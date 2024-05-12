package main

import "crypto/sha256"

func sha256_to_bytes(dat []byte) []byte {
	hash_32 := sha256.Sum256([]byte(dat))
	hash := hash_32[:]
	return hash
}
