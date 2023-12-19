package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declaring variables
	var int_value_first int = 1
	int_value_second := 1
	fmt.Println(reflect.TypeOf(int_value_first), reflect.TypeOf(int_value_second))

	// Declaring arrays
	var array_first [3]string
	array_first[0] = "3"
	var array_second = [3]string{"1", "2", "3"}
	fmt.Println(reflect.TypeOf(array_first[2]), array_first[0], array_second)

	// Copy arrays
	var array = [3]string{"fitst", "second", "third"}
	newArray := array
	
	array[2] = "test"
	fmt.Println("age", newArray)
}
