package lab0

import (
	"fmt"
	"crypto/sha3"
)

func sha3_256_hash(input1 string ,input2 string) {
	hash1 := sha3.Sum256([]byte(input1))
	hash2 := sha3.Sum256([]byte(input2))
	fmt.Println("Hash (SHA3-256):")
	fmt.Printf("Output A = %x\n", hash1)
	fmt.Printf("Output B = %x\n", hash2)
	if hash1 == hash2 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> Not Match!")
	}
}