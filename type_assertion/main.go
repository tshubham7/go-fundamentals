/*A type assertion provides access to an
interface value's underlying concrete value.
t := i.(T)
*/
package main

import (
	"fmt"
)

func main() {
	example()

	// use case, if you want to create a generic function
	useCase(3)       // passing interger value
	useCase("hello") // passing string value
}

func example() {
	// examples
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic if wrong type
	fmt.Println(f)
}

func useCase(i interface{}) {
	switch v := i.(type) {
	case int:
		// treat like integer
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// treat like string
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// handle if unknown
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
