package game_of_life_5

import "testing"

func TestIsAlive(t *testing.T) {
	actual:= IsAlive(1, []int {1,0,0,0,0,0,0,0})
	expected:= 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}