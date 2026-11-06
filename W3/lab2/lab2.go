package lab2

import (
	"W3/lab2/utils/crack"
	"bufio"
	"fmt"
	"os"
)

func Lab2() {
	target_hash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	filepath := "lab2/nord_vpn.txt"
	fmt.Println("======== SHA1 Password Cracking ========")
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