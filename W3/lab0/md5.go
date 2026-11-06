package lab0

import (
	"crypto/md5"
	"fmt"
)

func md5_hash(input1 string ,input2 string) {
	hash1 := md5.Sum([]byte(input1))
	hash2 := md5.Sum([]byte(input2))
	fmt.Println("Hash (MD5):")
	fmt.Printf("Output A = %x\n", hash1)
	fmt.Printf("Output B = %x\n", hash2)
	if hash1 == hash2 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> Not Match!")
	}
}