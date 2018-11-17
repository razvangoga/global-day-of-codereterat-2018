package game_of_life_2

import "fmt"

func Evolve(current [][]int) [][]int {

	result := make([][]int, len(current))
	fmt.Printf("%v", result)
	for i := 0; i < len(current); i++ {
		//result = append(result, []int{})

		for j:=0; j< len(current[i]); j++ {
			//fmt.Printf("%d %v", i, result)
			result[i] = append(result[i],0)
			result[i][j] = IsAlive(current[i][j], [8]int{1,1,1,1,1,1,1,1})
		}
	}

	return result
}
