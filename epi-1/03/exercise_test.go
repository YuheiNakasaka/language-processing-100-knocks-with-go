package exercise

import "testing"

func TestExercise(t *testing.T) {
	expected := []int{3, 1, 4, 1, 5, 1, 9, 2, 6, 1, 5, 3, 5, 8, 9, 7, 9, 1}
	res := exercise()
	for i := range res {
		if expected[i] != res[i] {
			t.Errorf("exercise() should return %v, but returned %v", expected[i], res[i])
		}
	}
}
