package exercise

import (
	"reflect"
	"testing"
)

func TestExercise(t *testing.T) {
	expected := map[string]int{"He": 2, "B": 5, "O": 8, "F": 9, "Ca": 20, "N": 7, "Na": 11, "Al": 13, "Cl": 17, "Ar": 18, "H": 1, "Be": 4, "C": 6, "Mi": 12, "Si": 14, "P": 15, "K": 19, "Li": 3, "Ne": 10, "S": 16}
	res := exercise()
	for k := range expected {
		if _, ok := res[k]; ok != true {
			keys := reflect.ValueOf(expected).MapKeys()
			t.Errorf("exercise() should returned a key from %v", keys)
		}

		if expected[k] != res[k] {
			t.Errorf("exercise() should return %v, but returned %v", expected[k], res[k])
		}
	}
}
