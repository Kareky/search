package main

import (
	"fmt"
	"github.com/Kareky/search/binary-search"
)

func main() {
	// Example usage of BinarySearchInt
	intArray := []int{1, 3, 5, 7, 9}
	targetInt := 9
	index, found := binarySearch.SearchInt(intArray, targetInt)
	if found {
		fmt.Printf("Found %d at index %d\n", targetInt, index)
	} else {
		fmt.Printf("%d not found in the array\n", targetInt)
	}

	// Example usage of BinarySearchString
	stringArray := []string{"apple", "banana", "cherry", "date", "fig"}
	targetString := "apple"
	index, found = binarySearch.SearchString(stringArray, targetString)
	if found {
		fmt.Printf("Found %s at index %d\n", targetString, index)
	} else {
		fmt.Printf("%s not found in the array\n", targetString)
	}

	// Example usage of BinarySearchBasicType
	floatArray := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	targetFloat := 2.2
	index, found = binarySearch.SearchBasicType(floatArray, targetFloat)
	if found {
		fmt.Printf("Found %f at index %d\n", targetFloat, index)
	} else {
		fmt.Printf("%f not found in the array\n", targetFloat)
	}
}