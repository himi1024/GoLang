package main

// Package import
import "fmt"

func main() {
	// variables
	var number uint
	var var1 string = "HelloWorld"
	sugar := "Sugar"

	// Arrays
	// Slices (Dynamic Arrays)
	// var anything [30]string
	anything := []string{}
	anything = append(anything, "QQQ")
	anything = append(anything, "TQQQ")

	fmt.Println(len(anything), "abcd")
	// Get user input by reference
	fmt.Println("Enter a positive number:")
	fmt.Scan(&number)

	// Validation
	isValidNumber := number <= 100

	// formatted Output: %v => variables
	if isValidNumber {
		fmt.Printf(" %v and %v and %v\n", var1, number, sugar)
		fmt.Printf("var1 is in %T type and number is in %T type\n", var1, number)
		fmt.Printf("The whole array: %v\n", anything)
	}

	// For-loop-1
	for i := 1; i < 3; i++ {
		fmt.Println("For-loop")
	}

	// For-loop-2
	for i, s := range anything {
		fmt.Println(i, s)
	}
}
