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
	fmt.Println("Copy arrays:", array, newArray)

	// Declaring Slices
	var slice []string
	var slice_first = []string{"1", "2", "3"}
	slice_second := []string{"1", "2", "3"}
	fmt.Println(slice, slice_first, slice_second)

	myslice := make([]string, 3) //a slice of 3 empty strings
	new_slice_first := append(slice_first, "test", "testd")
	fmt.Println(myslice, slice_first, new_slice_first)
	slice_first[1] = "4"
	fmt.Println(myslice, slice_first, new_slice_first)

	// some test
	array_test := [3]string{"First", "Second", "Third"}
	mySlice := array_test[:]
	array_test[1] = "4"
	fmt.Println(mySlice, array_test)
}
