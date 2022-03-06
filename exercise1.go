package main

import (
	"fmt"
	"os"
	"path"
)
func main() {
	// ---------------------------------------------------------
	// EXERCISE: Make It Blue
	//
	//  1. Change `color` variable's value to "blue"
	//
	//  2. Print it
	//
	// EXPECTED OUTPUT
	//  blue
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW:

	color = "blue"
	fmt.Println(color)
	// ---------------------------------------------------------
	// EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//
	//  2. Do not assign "dark green" to `color` directly
	//
	//     Instead, use the `color` variable again
	//     while assigning to `color`
	//
	//  3. Print it
	//
	// RESTRICTIONS
	//  WRONG ANSWER, DO NOT DO THIS:
	//  `color = "dark green"`
	//
	// HINT
	//  + operator can concatenate string values
	//
	//  Example:
	//
	//  "h" + "e" + "y" returns "hey"
	//
	// EXPECTED OUTPUT
	//  dark green
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	color = "green"

	// ADD YOUR CODE BELOW

	color = "dark " + color

	// UNCOMMENT THE CODE BELOW TO PRINT THE VARIABLE

	fmt.Println(color)
	// ---------------------------------------------------------
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//
	//  2. Print the `n` variable
	//
	// HINT
	//  Example: 3 * 2 = 6
	//  * is the product operator (it multiplies numbers)
	//
	// EXPECTED OUTPUT
	//  6.28
	// ---------------------------------------------------------
	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0
	n := 0.

	// ADD YOUR CODE BELOW

	n = 3.14 * 2 	

	fmt.Println(n)
	// ---------------------------------------------------------
	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//
	//  2. Assign the result to the `perimeter` variable
	//
	//  3. Print the `perimeter` variable
	//
	// HINT
	//  Formula = 2 times the width and height
	//
	// EXPECTED OUTPUT
	//  22
	//
	// BONUS
	//  Find more formulas here and calculate them in new programs
	//  https://www.mathsisfun.com/area.html
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	var (
		perimeter        int
		width, height = 5, 6
	)

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT

	// ADD YOUR CODE BELOW
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)
	// ---------------------------------------------------------
	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  go version 2
	// ---------------------------------------------------------
	// DO NOT TOUCH THIS
	var (
		lang    string
		version int
	)

	// ADD YOUR CODE BELOW
	lang, version = "go", 2

	// DO NOT TOUCH THIS
	fmt.Println(lang, "version", version)

	// ---------------------------------------------------------
	// EXERCISE: Multi Assign #2
	//
	//  1. Assign the correct values to the variables
	//     to match to the EXPECTED OUTPUT below
	//
	//  2. Print the variables
	//
	// HINT
	//  Use multiple Println calls to print each sentence.
	//
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees
	// ---------------------------------------------------------
	
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet, isTrue,	temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is",temp,"degrees")
	// ---------------------------------------------------------
	// EXERCISE: Multi Short Func
	//
	// 	1. Declare two variables using multiple short declaration syntax
	//
	//  2. Initialize the variables using `multi` function below
	//
	//  3. Discard the 1st variable's value in the declaration
	//
	//  4. Print only the 2nd variable
	//
	// NOTE
	//  You should use `multi` function
	//  while initializing the variables
	//
	// EXPECTED OUTPUT
	//  4
	// ---------------------------------------------------------

	// ADD YOUR DECLARATIONS HERE
	_,b := multi()

	// THEN UNCOMMENT THE CODE BELOW
	fmt.Println(b)

	// ---------------------------------------------------------
	// EXERCISE: Swapper
	//
	//  1. Change `color1` to "orange"
	//     and `color2` to "green" at the same time
	//
	//     (use multiple-assignment)
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  orange green
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	color1, color2 := "red", "blue"
	color1, color2 = "orange", "green"
	fmt.Println(color1, color2)
	// ---------------------------------------------------------
	// EXERCISE: Swapper #2
	//
	//  1. Swap the values of `red` and `blue` variables
	//
	//  2. Print them
	//
	// EXPECTED OUTPUT
	//  blue red
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)
	// ---------------------------------------------------------
	// EXERCISE: Discard The File
	//
	//  1. Print only the directory using `path.Split`
	//
	//  2. Discard the file part
	//
	// RESTRICTION
	//  Use short declaration
	//
	// EXPECTED OUTPUT
	//  secret/
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	folder, _:= path.Split("secret/file.txt")
	fmt.Println(folder)

	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix
	//
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  15.5
	// ---------------------------------------------------------
	
	a, b2 := 10, 5.5
	fmt.Println(float64(a) + b2)
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #2
	//
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  10.5
	// ---------------------------------------------------------
	a, b2 = 10, 5.5
	a = int(b2)
	fmt.Println(float64(a) + b2)
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #3
	//
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  5.5
	// ---------------------------------------------------------
	fmt.Println(float64(5.5))
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #4
	//
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  9.5
	// ---------------------------------------------------------
	age := 2
	fmt.Println(7.5 + float64(age))
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #5
	//
	//  Fix the code.
	//
	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	//
	// EXPECTED OUTPUT
	//  1127
	// ---------------------------------------------------------
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))

	// ---------------------------------------------------------
	// EXERCISE: Count the Arguments
	//
	//  Print the count of the command-line arguments
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 names.
	// ---------------------------------------------------------
	// UNCOMMENT & FIX THIS CODE
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)
	// ---------------------------------------------------------
	// EXERCISE: Print Your Name
	//
	//  Get it from the first command-line argument
	//
	// INPUT
	//  Call the program using your name
	//
	// EXPECTED OUTPUT
	//  It should print your name
	//
	// EXAMPLE
	//  go run main.go inanc
	//
	//    inanc
	//
	// BONUS: Make the output like this:
	//
	//  go run main.go inanc
	//    Hi inanc
	//    How are you?
	// ---------------------------------------------------------
	fmt.Println("Hi",os.Args[4])
	fmt.Println("How are you?")
	// ---------------------------------------------------------
	// EXERCISE: Greet More People
	//
	// RESTRICTIONS
	//  1. Be sure to match the expected output below
	//  2. Store the length of the arguments in a variable
	//  3. Store all the arguments in variables as well
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 people!
	//  Hello great bilbo !
	//  Hello great balbo !
	//  Hello great bungo !
	//  Nice to meet you all.
	// ---------------------------------------------------------
	// TYPE YOUR CODE HERE
	fmt.Printf("There are %d people.\n", count)
	fmt.Printf("Hello great %s!\n", os.Args[1])
	fmt.Printf("Hello great %s!\n", os.Args[2])
	fmt.Printf("Hello great %s!\n", os.Args[3])
	fmt.Printf("Hello great %s!\n", os.Args[4])
	fmt.Printf("Nice to meet you all.\n")
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.

}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}