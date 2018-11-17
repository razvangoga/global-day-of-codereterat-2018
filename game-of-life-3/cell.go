package game_of_life_3

import "fmt"

func CellEvolve(alive int, neighbours [8]int) int {

	sum := sumOfNeighbours(neighbours);

	if alive == 0 {
		if sum == 3 {
			return 1
		}
		return 0
	}

	if sum < 2 {
		fmt.Print(sum)
		return 0
	}
	if sum == 2 || sum == 3 {
		return 1
	}

	return 0

}

func sumOfNeighbours(neighbours [8]int) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += neighbours[i]
	}
	return result
}
