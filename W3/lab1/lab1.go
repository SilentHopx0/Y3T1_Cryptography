package lab1

import (
	"W3/lab1/utils/crack"
	"bufio"
	"fmt"
	"os"
)

func Lab1() {
	target_hash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	filepath := "lab1/nord_vpn.txt"
	fmt.Println("======== MD5 Password Cracking ========")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		hash := crack.Md5_hash(word)
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