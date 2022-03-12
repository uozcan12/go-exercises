package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func main() {
	// ---------------------------------------------------------
	// STORY
	//
	//  Your boss wants you to create a program that will check
	//  whether a person can watch a particular movie or not.
	//
	//  Imagine that another program provides the age to your
	//  program. Depending on what you return, the other program
	//  will issue the tickets to the person automatically.
	//
	// EXERCISE: Movie Ratings
	//
	//  1. Get the age from the command-line.
	//
	//  2. Return one of the following messages if the age is:
	//     -> Above 17         : "R-Rated"
	//     -> Between 13 and 17: "PG-13"
	//     -> Below 13         : "PG-Rated"
	//
	// RESTRICTIONS:
	//  1. If age data is wrong or absent let the user know.
	//  2. Do not accept negative age.
	//
	// BONUS:
	//  1. BONUS: Use if statements only twice throughout your program.
	//  2. BONUS: Use an if statement only once.
	//
	// EXPECTED OUTPUT
	//  go run main.go 18
	//    R-Rated
	//
	//  go run main.go 17
	//    PG-13
	//
	//  go run main.go 12
	//    PG-Rated
	//
	//  go run main.go
	//    Requires age
	//
	//  go run main.go -5
	//    Wrong age: "-5"
	// ---------------------------------------------------------
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
	
	// 🛑 DON'T DO THIS:
	// It's hard to read.
	// It's just an exercise.

	// if len(os.Args) != 2 {
	// 	fmt.Println("Requires age")
	// 	return
	// } else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
	// 	fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
	// 	return
	// } else if age > 17 {
	// 	fmt.Println("R-Rated")
	// } else if age >= 13 && age <= 17 {
	// 	fmt.Println("PG-13")
	// } else if age < 13 {
	// 	fmt.Println("PG-Rated")
	// }

	// ---------------------------------------------------------
	// EXERCISE: Odd or Even
	//
	//  1. Get a number from the command-line.
	//
	//  2. Find whether the number is odd, even and divisible by 8.
	//
	// RESTRICTION
	//  Handle the error. If the number is not a valid number,
	//  or it's not provided, let the user know.
	//
	// EXPECTED OUTPUT
	//  go run main.go 16
	//    16 is an even number and it's divisible by 8
	//
	//  go run main.go 4
	//    4 is an even number
	//
	//  go run main.go 3
	//    3 is an odd number
	//
	//  go run main.go
	//    Pick a number
	//
	//  go run main.go ABC
	//    "ABC" is not a number
	// ---------------------------------------------------------
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}

	// ---------------------------------------------------------
	// EXERCISE: Leap Year
	//
	//  Find out whether a given year is a leap year or not.
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Give me a year number
	//
	//  go run main.go eighties
	//    "eighties" is not a valid year.
	//
	//  go run main.go 2018
	//    2018 is not a leap year.
	//
	//  go run main.go 2100
	//    2100 is not a leap year.
	//
	//  go run main.go 2019
	//    2019 is not a leap year.
	//
	//  go run main.go 2020
	//    2020 is a leap year.
	//
	//  go run main.go 2024
	//    2024 is a leap year.
	// ---------------------------------------------------------
	
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", os.Args[1])
		return
	}

	var leap bool
	var year = n
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}

	// ---------------------------------------------------------
	// EXERCISE: Simplify the Leap Year
	//
	//  1. Look at the solution of "the previous exercise".
	//
	//  2. And simplify the code (especially the if statements!).
	//
	// EXPECTED OUTPUT
	//  It's the same as the previous exercise.
	// ---------------------------------------------------------
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number.")
		return
	}

	n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", os.Args[1])
		return
	}

	if (n % 400 == 0) && (n % 100 == 0) {
		fmt.Printf("%d is a leap year\n", n)
	} else if(n % 4 ==0) && (n % 100 != 0) { // not divided by 100 means not a century year year divided by 4 is a leap year
		fmt.Printf("%d is a leap year\n", n)
	} else { // if not divided by both 400 (century year) and 4 (not century year) year is not leap year
		fmt.Printf("%d is not a leap year\n", n)
	}
}