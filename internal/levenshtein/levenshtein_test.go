package levenshtein_test

import (
	"testing"
	"github.com/Kareky/search/internal/levenshtein"
)

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		s1	   	string
		s2	   	string
		expect	int
	} {
		{"kitten", "sitting", 3},
		{"flaw", "lawn", 2},
		{"intention", "execution", 5},
	}

	for _, test := range tests {
		distance := levenshtein.CalculateDistance(test.s1, test.s2)
		if distance != test.expect {
			t.Errorf("Expected distance between %s and %s to be %d, but got %d", test.s1, test.s2, test.expect, distance)
		}
	}
}