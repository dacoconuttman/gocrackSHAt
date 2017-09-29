package utils

import (
	"crypto/sha256"
	"fmt"
)

type Tuple struct {
	Plaintext string
	Hash      string
}

//Hash hashes a string, returns a 32 byte long byte array
func Hash(plaintext string) (sha string) {
	return fmt.Sprintf("%s", sha256.Sum256([]byte(plaintext)))
}

//HashConc hashes a string, then puts the result into a channel
func HashConc(plaintext string, c chan Tuple) {
	c <- Tuple{Plaintext: plaintext, Hash: fmt.Sprintf("%s", sha256.Sum256([]byte(plaintext)))}
}
