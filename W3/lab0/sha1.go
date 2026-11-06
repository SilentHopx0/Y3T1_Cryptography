package lab0

import (
	"fmt"
	"crypto/sha1"
)

func sha1_hash(input1 string ,input2 string) {
	hash1 := sha1.Sum([]byte(input1))
	hash2 := sha1.Sum([]byte(input2))
	fmt.Println("Hash (SHA1):")
	fmt.Printf("Output A = %x\n", hash1)
	fmt.Printf("Output B = %x\n", hash2)
	if hash1 == hash2 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> Not Match!")
	}
}