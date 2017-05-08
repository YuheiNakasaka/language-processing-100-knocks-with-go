package reverse

import "testing"

func TestReverse(t *testing.T) {
	if reverse() != "desserts" {
		t.Fatal("reverse() should return desserts")
	}
}
