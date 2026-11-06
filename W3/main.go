package main

import (
	"W3/lab0"
	"W3/lab1"
	"W3/lab2"
	"W3/lab3"
	"fmt"
)

func main() {
	var choice string
	fmt.Println("=== W3 Lab Exercises ===")
	for i := 0; i <= 3; i++ {
		fmt.Printf("%d. Exercise %d\n", i, i)
	}
	fmt.Print("Choose your exercise (0-3): ")
	fmt.Scan(&choice)

	switch choice {
	case "0":
		lab0.Lab0()
	case "1":
		lab1.Lab1()
	case "2":
		lab2.Lab2()
	case "3":
		lab3.Lab3()
	default:
		fmt.Println("Invalid Option.")
	}
}