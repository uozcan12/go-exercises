package main
import (
	"fmt"
	"path"
	"os"
)
func main() {
	//var file string
	//_, file = path.Split("css/main.css")
	_, file := path.Split("css/main.css") //short declaration
	fmt.Println("Hello")
	fmt.Println("file:", file)

	var score int // already score 0
	width, height := 100,50 //short declaration
	width, color := 50, "red" //short declaration
	fmt.Println("score:", score)
	fmt.Println("width:", width)
	fmt.Println("height:", height)
	fmt.Println("color:", color)
	
	//Type conversion

	speed := 100
	force := 2.5
	speed1 := speed * int(force)
	speed2 := int(float64(speed) * force)

	fmt.Println("speed1:", speed1)
	fmt.Println("speed2:", speed2)

	// Intro to slices
	// Args's type is a "slice of strings"
	
	fmt.Println("%#v\n:", os.Args)



}