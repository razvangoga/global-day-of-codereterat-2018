package game_of_life_5

import "testing"

func TestIsAlive(t *testing.T) {
	actual:= IsAlive(1, 1)
	expected:= 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_live_2(t *testing.T) {
	actual:= IsAlive(1, 2)
	expected:= 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_3(t *testing.T) {
	actual:= IsAlive(1, 3)
	expected:= 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_more(t *testing.T) {
	actual:= IsAlive(1, 4)
	expected:= 0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}


func TestIsAlive_revive(t *testing.T) {
	actual:= IsAlive(0, 3)
	expected:= 1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}







