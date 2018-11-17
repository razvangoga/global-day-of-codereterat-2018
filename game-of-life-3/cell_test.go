package game_of_life_3

import "testing"


func TestCellEvolve_2_neighbours(t *testing.T) {
	actual := CellEvolve(1, [8]int{1, 1,0 , 0, 0, 0, 0,0})
	expected :=1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestCellEvolve_underpopulation(t *testing.T) {
	actual := CellEvolve(1, [8]int{1, 0, 0, 0, 0,0,0,0})
	expected :=0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestCellEvolve_revive(t *testing.T) {
	actual := CellEvolve(0, [8]int{1, 1, 1, 0, 0, 0, 0,0})
	expected :=1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

