package lab0

import (
	"fmt"
)

func Lab0() {
	var input1 string
	var input2 string
	fmt.Println("======== Sin Hong Hashing Program ========")
	fmt.Print("Please input value 1: ")
	fmt.Scan(&input1)
	fmt.Print("Please input value 2: ")
	fmt.Scan(&input2)
	md5_hash(input1, input2)
	sha1_hash(input1, input2)
	sha256_hash(input1, input2)
	sha512_hash(input1, input2)
	sha3_256_hash(input1, input2)
}