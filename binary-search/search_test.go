package binarySearch_test

import (
	"testing"
	"github.com/Kareky/search/binary-search"
)

func TestBinarySearchInt(t *testing.T) {
	intArray := []int{-8, 0, 11, 12, 22, 25, 64}
	target := 22
	expectedIndex := 4
	index, found := binarySearch.SearchInt(intArray, target)
	if !found {
		t.Errorf("Expected to find %d but it was not found", target)
	}
	if index != expectedIndex {
		t.Errorf("Expected index %d but got %d", expectedIndex, index)
	}
}

func TestBinarySearchString(t *testing.T) {
	stringArray := []string{"apple", "banana", "cherry", "date"}
	target := "cherry"
	expectedIndex := 2
	index, found := binarySearch.SearchString(stringArray, target)
	if !found {
		t.Errorf("Expected to find %s but it was not found", target)
	}
	if index != expectedIndex {
		t.Errorf("Expected index %d but got %d", expectedIndex, index)
	}
}

func TestBinarySearchBasicType(t *testing.T) {
	floatArray := []float64{-8.5, 0.0, 11.0, 12.1, 22.2, 25.3, 64.5}
	target := 12.1
	expectedIndex := 3
	index, found := binarySearch.SearchBasicType(floatArray, target)
	if !found {
		t.Errorf("Expected to find %f but it was not found", target)
	}
	if index != expectedIndex {
		t.Errorf("Expected index %d but got %d", expectedIndex, index)
	}
}

func TestBinarySearchSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	target := Person{Name: "Alice", Age: 30}
	expectedIndex := 1
	index, err := binarySearch.SearchSlice(personArray, target, func(p1, p2 Person) int {
		if p1.Age < p2.Age {
			return -1
		} else if p1.Age > p2.Age {
			return 1
		} else {
			return 0
		}
	})
	if err != nil {
		t.Errorf("Unexpected error occurred: %v", err)
	}
	if index != expectedIndex {
		t.Errorf("Expected index %d but got %d", expectedIndex, index)
	}
}