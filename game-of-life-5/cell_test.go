package game_of_life_5

import "testing"

func TestIsAlive(t *testing.T) {
	actual:= IsAlive(1, []int {1,0,0,0,0,0,0,0})
	expected:= 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_live_2(t *testing.T) {
	actual:= IsAlive(1, []int {1,1,0,0,0,0,0,0})
	expected:= 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_3(t *testing.T) {
	actual:= IsAlive(1, []int {1,1,1,0,0,0,0,0})
	expected:= 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_more(t *testing.T) {
	actual:= IsAlive(1, []int {1,1,1,1,0,0,0,0})
	expected:= 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_revive(t *testing.T) {
	actual:= IsAlive(0, []int {1,1,1,0,0,0,0,0})
	expected:= 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}







