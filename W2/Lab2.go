package main

import (
	"fmt"
)

func exercise1() {
	var num1 int
	var num2 int
	var operator string
	var result int
	fmt.Print("Enter first number (Main Number): ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Choose your operator (=, +=, -=, *=, /=, %=): ")
	fmt.Scan(&operator)
	switch operator {
		case "=":
			num1 = num2
			result = num1
		case "+=":
			num1 += num2
			result = num1
		case "-=":
			num1 -= num2
			result = num1
		case "*=":
			num1 *= num2
			result = num1
		case "/=":
			if num2 != 0 {
				num1 /= num2
				result = num1
			} else {
				fmt.Println("Error: Cannot devide by 0")
			}
		case "%=":
			if num2 != 0 {
				num1 %= num2
				result = num1
			} else {
				fmt.Println("Error: Cannot modulo by 0")
			}
		default:
			fmt.Print("Invalid Operator.")
	}
	fmt.Println("Result: ", result)
} 

func exercise2() {
	var num1 int
	var num2 int
	var operator string
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Choose your operator (&&, ||, !): ")
	fmt.Scan(&operator)
	switch operator {
	case "&&":
		var result bool
		if (num1 > 0 && num2 > 0) {
			result = true
		} else {
			result = false
		}
		fmt.Println("Result: ", result)
	case "||":
		var result bool
		if (num1 > num2 || num2 > num1) {
			result = true
		} else {
			result = false
		}
		fmt.Println("Result: ", result)
	case "!":
		var result bool
		if (num1 != num2) {
			result = true
		} else {
			result = false
		}
		fmt.Println("Result: ", result)
	}
}



func exercise3() {
	var num1 int
	var num2 int
	var operator string
	var result int
	var shift int
	fmt.Print("Enter first number (Main Number): ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Choose your operator (AND, OR, XOR, NOT, ls(left shift), rs(right shift)): ")
	fmt.Scan(&operator)
	switch operator {
		case "AND":
			result = num1 & num2
		case "OR":
			result = num1 | num2
		case "XOR":
			result = num1 ^ num2
		case "NOT":
			result = ^num1
		case "ls":
			fmt.Println("Enter the value of left shift: ")
			fmt.Scan(&shift)
			result = num1 << shift
		case "rs":
			fmt.Println("Enter the value of right shift: ")
			fmt.Scan(&shift)
			result = num1 >> shift
		default:
			fmt.Print("Invalid Operator.")
	}
	fmt.Printf("Result: %8b\n", result)
}

func main() {
	var choice string
	fmt.Println("W2 Lab Exercise:\n")
	for i := 1; i < 7; i++ {
		fmt.Printf("%v. Exercise %v\n" , i, i)
	}
	fmt.Print("Choose your exercise: ")
	fmt.Scan(&choice)
	switch choice {
		case "1":
			exercise1()
		case "2":
			exercise2()
		case "3":
			exercise3()
		case "4":
			exercise1()
		default:
			fmt.Println("Invalid Operation.")
	}
	
}
