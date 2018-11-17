package game_of_life_4

import "testing"

func TestIsAlive_Rule1(t *testing.T) {
	actual  := IsAlive(1, [8]int {1, 0,0,0,0,0,0,0})
	expected := 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestIsAlive_Rule2(t *testing.T) {
	actual  := IsAlive(1, [8]int {1, 1,1,0,0,0,0,0})
	expected := 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestIsAlive_Rule3(t *testing.T) {
	actual  := IsAlive(1, [8]int {1, 1,1,1,1,1,1,1})
	expected := 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestIsAlive_Rule4(t *testing.T) {
	actual  := IsAlive(0, [8]int {1, 1,1,0,0,0,0,0})
	expected := 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}