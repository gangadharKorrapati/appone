package main

import (
	"fmt"
	"strconv"
)

var n float32 = 42

//lower case variables are scoped to packages
//upper case letters are exposed to outside package
var (
	count int = 0
	score int = 12
)

func main() {
	fmt.Println("Hello world")
	var i int
	i = 41
	fmt.Printf("%v %T\n", i, i)
	var j int = 42
	fmt.Printf("%v %T\n", j, j)
	var k = 43
	fmt.Printf("%v %T\n", k, k)
	l := 42
	fmt.Printf("%v %T\n", l, l)
	m := 42.
	fmt.Printf("%v %T\n", m, m)
	//no decleration for float 32  with :=
	//can't use the := at package level
	fmt.Printf("%v %T\n", n, n)

	fmt.Printf("count %v score %v\n", count, score)

	//it is a crime to create a variable and not use in go
	//block scoped variables are not accessable outside package

	//acronyms are wrotten in uppercases
	var theURL string = "https://gangadhar.me"

	fmt.Println(theURL)

	var x int = 7
	var y float32 = float32(x)
	fmt.Printf("%v %T %v %T\n", x, x, y, y)
	y = 22.7
	x = int(y)
	//info lost .7
	fmt.Printf("%v %T %v %T\n", x, x, y, y)
	// string is alias for bytes
	fmt.Println(string(97))
	// to convert numbert to string use Itoa
	fmt.Println(strconv.Itoa(97))

	//shadowing the package level variables
	fmt.Printf("count %v\n", count)
	var count int = 32
	fmt.Printf("count %v\n", count)
}
