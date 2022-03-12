package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
	// ---------------------------------------------------------
	// EXERCISE: Only Evens
	//
	//  1. Extend the "Sum up to N" exercise
	//  2. Sum only the even numbers
	//
	// RESTRICTIONS
	//  Skip odd numbers using the `continue` statement
	//
	// EXPECTED OUTPUT
	//  Let's suppose that the user runs it like this:
	//
	//    go run main.go 1 10
	//
	//  Then it should print:
	//
	//    2 + 4 + 6 + 8 + 10 = 30
	// ---------------------------------------------------------
	sum=0
	if(max%2!=0){
		max = max-1
	}
	for i := min; i <= max; i++ {
		if i%2==0{
			sum += i
			fmt.Print(i)
			if i < max {
				fmt.Print(" + ")
			}
		}

	}
	fmt.Printf(" = %d\n", sum)
	// var sum int
	// for i := min; i <= max; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	sum += i

	// 	fmt.Print(i)
	// 	if i < max-1 {
	// 		fmt.Print(" + ")
	// 	}
	// }
	// fmt.Printf(" = %d\n", sum)
	// ---------------------------------------------------------
	// EXERCISE: Break Up
	//
	//  1. Extend the "Only Evens" exercise
	//  2. This time, use an infinite loop.
	//  3. Break the loop when it reaches to the `max`.
	//
	// RESTRICTIONS
	//  You should use the `break` statement once.
	//
	// HINT
	//  Do not forget incrementing the `i` before the `continue`
	//  statement and at the end of the loop.
	//
	// EXPECTED OUTPUT
	//    go run main.go 1 10
	//    2 + 4 + 6 + 8 + 10 = 30
	// ---------------------------------------------------------
	sum=0
	var (
		i   = min
	)

	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}

		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)
	// ---------------------------------------------------------
	// EXERCISE: Infinite Kill
	//
	//  1. Create an infinite loop
	//  2. On each step of the loop print a random character.
	//  3. And, sleep for 250 milliseconds.
	//  4. Run the program and wait a couple of seconds
	//     then kill it using CTRL+C keys
	//
	// RESTRICTIONS
	//  1. Print one of those characters randomly: \ / | -
	//  2. Before printing a character print also this
	//     escape sequence: \r
	//
	//     Like this: "\r/", or this: "\r|", and so on...
	//
	//  3. NOTE : If you're using Go Playground, use "\f" instead of "\r"
	//
	// HINTS
	//  1. Use `time.Sleep` to sleep.
	//  2. You should pass a `time.Duration` value to it.
	//  3. Check out the Go online documentation for the
	//     `Millisecond` constant.
	//  4. When printing, do not use a newline! Or a Println!
	//     Use Print or Printf instead.
	//
	// NOTE
	//  If this exercise is hard for you now, wait until the
	//  lucky number lecture. Even then so, then just skip it.
	//
	//  You can return back to it afterwards.
	//
	// EXPECTED OUTPUT
	//  - The program should display the following messages in any order.
	//  - And, the first character (\, -, /, or |) should be randomly
	//    displayed.
	//  - \r or \f sequence will clear the line.
	//
	//  \ Please Wait. Processing....
	//  - Please Wait. Processing....
	//  / Please Wait. Processing....
	//  | Please Wait. Processing....
	// ---------------------------------------------------------
	for {
		var c string

		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}