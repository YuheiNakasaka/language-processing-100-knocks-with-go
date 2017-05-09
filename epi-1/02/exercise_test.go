package exercise

import "testing"

func TestExercise(t *testing.T) {
	expected := "パタトクカシーー"
	if res := exercise(); res != expected {
		t.Errorf("exercise() should return %v, but returned %v", expected, res)
	}
}
