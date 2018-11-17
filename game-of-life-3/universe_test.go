package game_of_life_3

import (
	"reflect"
	"testing"
)

func TestEvolve_1(t *testing.T) {
	actual := Evolve([][]int{ {0,1,0}, {0,1,0}, {0,1,0}})
	expected := [][]int{ {0,0,0}, {1,0,1}, {0,0,0}}

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}

func TestEvolve_2(t *testing.T) {
	actual := Evolve([][]int{ {0,1,0}, {1,1,1}, {0,1,1}})
	expected := [][]int{ {1,1,1}, {1,0,1}, {1,1,1}}

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}
