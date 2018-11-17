package game_of_life_4

func Evolve(current [][]int) [][]int {
	result := make([][]int, len(current))

	for i := 0; i < len(current); i++ {
		result[i] = make([]int, len(current[i]))

		for j := 0; j < len(current[i]); j++ {
			neighbours := extractNeighbors(i, j, current)
			result[i][j] = IsAlive(current[i][j], neighbours)
		}
	}
	return result
}

func extractNeighbors(x, y int, current [][]int) [8]int {
	res := [8]int{}

	index := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if x+i >= 0 && y+j >= 0 && x+i < len(current) && y+j < len(current[x]) {
				res[index] = current[x+i][y+j]
			}

			index++
		}
	}
	return res
}
