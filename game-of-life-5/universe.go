package game_of_life_5

func Evolve(current [][]int) [][]int {
	result := make([][]int, len(current))

	for i := 0; i < len(current); i++ {
		result[i] = make([]int, len(current[i]))

		for j := 0; j < len(current[i]); j++ {
			result[i][j] = IsAlive(current[i][j], getNeighbours(current, i, j))
		}
	}
	return result
}

func getNeighbours(universe [][]int, x int, y int) []int {

	result := make([]int, 8)
	idx := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			xx := x + i
			yy := y + j

			if xx < 0 || yy < 0 || xx >= len(universe) || yy >= len(universe) {
				continue
			}

			result[idx] = universe[xx][yy]
			idx++
		}
	}

	return result
}
