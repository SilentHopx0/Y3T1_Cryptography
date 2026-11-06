package lab3

import (
	"W3/lab3/utils/crack"
	"bufio"
	"fmt"
	"os"
)

func Lab3() {
	target_hash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	filepath := "lab3/nord_vpn.txt"
	fmt.Println("======== SHA512 Password Cracking ========")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		hash := crack.Sha1_hash(word)
		if hash == target_hash {
			fmt.Printf("Target Hash: %s\n", hash)
			fmt.Printf("Password Found: %s\n", word)
			break
		}
	}

	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file scanner:", err)
    }
}