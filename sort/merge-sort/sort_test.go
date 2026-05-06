package mergeSort_test

import (
	"testing"
	"github.com/Kareky/search/sort/merge-sort"
)

func TestMergeSortInt(t *testing.T) {
	intArray := []int{64, 25, 12, 22, 11, -8, 0}
	expected := []int{-8, 0, 11, 12, 22, 25, 64}
	mergeSort.SortInt(intArray)
	for i := range intArray {
		if intArray[i] != expected[i] {
			t.Errorf("Expected %d but got %d at index %d", expected[i], intArray[i], i)
		}
	}
}

func TestMergeSortString(t *testing.T) {
	stringArray := []string{"banana", "apple", "cherry", "date"}
	expected := []string{"apple", "banana", "cherry", "date"}
	mergeSort.SortString(stringArray)
	for i := range stringArray {
		if stringArray[i] != expected[i] {
			t.Errorf("Expected %s but got %s at index %d", expected[i], stringArray[i], i)
		}
	}
}

func TestMergeSortBasicType(t *testing.T) {
	floatArray := []float64{64.5, 25.3, 12.1, 22.2, 11.0, -8.5, 0.0}
	expected := []float64{-8.5, 0.0, 11.0, 12.1, 22.2, 25.3, 64.5}
	mergeSort.SortBasicType(floatArray)
	for i := range floatArray {
		if floatArray[i] != expected[i] {
			t.Errorf("Expected %f but got %f at index %d", expected[i], floatArray[i], i)
		}
	}
}

func TestMergeSortSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	expected := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	mergeSort.SortSlice(personArray, func(p1, p2 Person) int {
		if p1.Age < p2.Age {
			return -1
		} else if p1.Age > p2.Age {
			return 1
		} else {
			return 0
		}
	})
	for i := range personArray {
		if personArray[i] != expected[i] {
			t.Errorf("Expected %v but got %v at index %d", expected[i], personArray[i], i)
		}
	}
}