package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Lab #1: Assignment Operators

func exercise1() {
	var a, b int
	fmt.Print("Enter first number (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second number (b): ")
	fmt.Scan(&b)

	orig := a
	fmt.Println("\n=== Demonstrating Assignment Operators ===")

	temp := orig
	temp = b
	fmt.Printf("= : a = b => a = %d\n", temp)

	temp = orig
	temp += b
	fmt.Printf("+= : a += b => a = %d\n", temp)

	temp = orig
	temp -= b
	fmt.Printf("-= : a -= b => a = %d\n", temp)

	temp = orig
	temp *= b
	fmt.Printf("*= : a *= b => a = %d\n", temp)

	if b != 0 {
		temp = orig
		temp /= b
		fmt.Printf("/=  : a /= b => a = %d\n", temp)

		temp = orig
		temp %= b
		fmt.Printf("%%=  : a %%= b => a = %d\n", temp)
	} else {
		fmt.Println("/=  : Error: Cannot divide by 0")
		fmt.Println("%=  : Error: Cannot modulo by 0")
	}
}

// Lab #2: Logical Operators (&&, ||, !)

func exercise2() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Logical expressions as specified in Lab #2
	bothPositive := (num1 > 0) && (num2 > 0)
	oneGreater := (num1 > num2) || (num2 > num1)
	notEqual := !(num1 == num2)

	fmt.Println("\n=== Logical Operators Evaluation ===")
	fmt.Printf("Both positive (&&): (%d > 0 && %d > 0) => %t\n", num1, num2, bothPositive)
	fmt.Printf("One greater than other (||): (%d > %d || %d > %d) => %t\n", num1, num2, num2, num1, oneGreater)
	fmt.Printf("Not equal (!): !(%d == %d) => %t\n", num1, num2, notEqual)
}

// Lab #3: Bitwise and Assignment Operators

func myAND(a, b uint8) uint8 { return a & b }
func myOR(a, b uint8) uint8 { return a | b }
func myXOR(a, b uint8) uint8 { return a ^ b }
func myNOT(a uint8) uint8 { return ^a }

func exercise3() {
	var a, b uint8
	fmt.Print("Enter first number a (0-255): ")
	fmt.Scan(&a)
	fmt.Print("Enter second number b (0-255): ")
	fmt.Scan(&b)

	fmt.Printf("\na in binary: %08b (%d)\n", a, a)
	fmt.Printf("b in binary: %08b (%d)\n", b, b)
	fmt.Println("==============================")
	fmt.Printf("myAND(a, b): %08b (%d)\n", myAND(a, b), myAND(a, b))
	fmt.Printf("myOR(a, b): %08b (%d)\n", myOR(a, b), myOR(a, b))
	fmt.Printf("myXOR(a, b): %08b (%d)\n", myXOR(a, b), myXOR(a, b))
	fmt.Printf("myNOT(a): %08b (%d)\n", myNOT(a), myNOT(a))
	fmt.Printf("Left shift (a << 1): %08b (%d)\n", a<<1, a<<1)
	fmt.Printf("Right shift (a >> 1): %08b (%d)\n", a>>1, a>>1)
	temp := a
	temp &= b
	fmt.Printf("Assignment (a &= b): %08b (%d)\n", temp, temp)
}

// Lab #4: Mini Calculator

func exercise4() {
	for {
		fmt.Println("\n=== Mini Calculator ===")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting Mini Calculator...")
			break
		}

		if choice < 1 || choice > 6 {
			fmt.Println("Invalid choice! Choose 1-6.")
			continue
		}

		var a, b float64
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Printf("Result: %g + %g = %g\n", a, b, a+b)
		case 2:
			fmt.Printf("Result: %g - %g = %g\n", a, b, a-b)
		case 3:
			fmt.Printf("Result: %g * %g = %g\n", a, b, a*b)
		case 4:
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Printf("Result: %g / %g = %g\n", a, b, a/b)
			}
		case 5:
			if int(b) == 0 {
				fmt.Println("Error: modulo by zero")
			} else {
				fmt.Printf("Result: %d %% %d = %d\n", int(a), int(b), int(a)%int(b))
			}
		}
	}
}

// Lab #5: Binary, Hex, and Base64 Encoding

func exercise5() {
	var input string
	fmt.Print("Enter text to encode: ")
	fmt.Scan(&input)

	bytes := []byte(input)

	fmt.Println("\n=== Encoding Output ===")
	fmt.Printf("Original Text: %s\n", input)
	fmt.Print("Binary: ")
	for _, b := range bytes {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()
	fmt.Printf("Hexadecimal: %s\n", hex.EncodeToString(bytes))
	fmt.Printf("Base64: %s\n", base64.StdEncoding.EncodeToString(bytes))
}

// Lab #6: XOR Encryption / Decryption

func xorEncrypt(text string, key byte) string {
	bytes := []byte(text)
	result := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		result[i] = bytes[i] ^ key
	}
	return string(result)
}

func exercise6() {
	var message string
	fmt.Print("Enter plaintext message: ")
	fmt.Scan(&message)

	var keyStr string
	fmt.Print("Enter a single-character key: ")
	fmt.Scan(&keyStr)
	key := byte(keyStr[0])

	ciphertext := xorEncrypt(message, key)
	decrypted := xorEncrypt(ciphertext, key)

	fmt.Println("\n=== XOR Encryption/Decryption ===")
	fmt.Printf("Original Message: %s\n", message)
	fmt.Printf("Key: %c (0x%02x)\n", key, key)
	fmt.Printf("Ciphertext (Hex): %s\n", hex.EncodeToString([]byte(ciphertext)))
	fmt.Printf("Ciphertext (Raw): %s\n", ciphertext)
	fmt.Printf("Decrypted Message: %s\n", decrypted)
}

// Main Menu

func main() {
	var choice string
	fmt.Println("=== W2 Lab Exercises ===")
	for i := 1; i <= 6; i++ {
		fmt.Printf("%d. Exercise %d\n", i, i)
	}
	fmt.Print("Choose your exercise (1-6): ")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		exercise1()
	case "2":
		exercise2()
	case "3":
		exercise3()
	case "4":
		exercise4()
	case "5":
		exercise5()
	case "6":
		exercise6()
	default:
		fmt.Println("Invalid Option.")
	}
}
