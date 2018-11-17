package game_of_life_3

import "fmt"

func Evolve(current [][]int) [][]int {
	res := make([][]int, len(current))

	for i, _ := range current {
		for j := 0; j < len(current[i]); j++ {
			res[i] = append(res[i], 0)
		}
	}

	fmt.Println(res)
	for i := 0; i < len(current); i++ {
		for j := 0 ; j < len(current[i]); j++ {
			res[i][j] = CellEvolve(current[i][j], neighbours(i, j, current))
			fmt.Println( neighbours(i, j, current), "current", current[i][j], "next", res[i][j])
		}
	}

	return res
}

func neighbours(i, j int, current [][]int) [8]int {
	res := [8]int{}

	if i > 0 && j > 0 {
		res[0] = current[i - 1][j - 1]
	}

	if j > 0 {
		res[1] = current[i][j - 1]
	}

	if i > 0 {
		res[2] = current[i - 1][j]
	}

	if i < len(current) -1 && j < len(current) -1 {
		res[3] = current[i + 1][j + 1]
	}

	if i < len(current) -1  {
		res[4] = current[i + 1][j]
	}

	if  j < len(current) -1 {
		res[5] = current[i][j + 1]
	}

	if i > 0 && j < len(current) - 1 {
		res[6] = current[i -1][j + 1]
	}

	if i < len(current) - 1 && j > 0 {
		res[6] = current[i + 1][j - 1]
	}

	return res
}