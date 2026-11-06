package lab0

import (
	"fmt"
	"crypto/sha512"
)

func sha512_hash(input1 string ,input2 string) {
	hash1 := sha512.Sum512([]byte(input1))
	hash2 := sha512.Sum512([]byte(input2))
	fmt.Println("Hash (SHA512):")
	fmt.Printf("Output A = %x\n", hash1)
	fmt.Printf("Output B = %x\n", hash2)
	if hash1 == hash2 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> Not Match!")
	}
}