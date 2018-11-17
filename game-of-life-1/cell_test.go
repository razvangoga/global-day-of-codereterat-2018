package gol1

import "testing"

func TestIsCellAliveRule1(t *testing.T) {
	actual := IsCellAlive(1, [8]int{1, 0,0,0,0,0,0,0})
	expected :=0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestIsCellAliveRule2(t *testing.T) {
	actual := IsCellAlive(1, [8]int{1, 1,1,0,0,0,0,0})
	expected :=1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}
func TestIsCellAliveRule3(t *testing.T) {
	actual := IsCellAlive(1, [8]int{1, 1,1,1,1,1,1,1})
	expected :=0

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}
func TestIsCellAliveRule4(t *testing.T) {
	actual := IsCellAlive(0, [8]int{1, 1,1,0,0,0,0,0})
	expected :=1

	if actual != expected {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}