package game_of_life_5

func IsAlive(cell int, neighbours []int)int {

	sum:=0

	for i:=0;i<len(neighbours); i++  {
		sum+= neighbours[i]
	}

	if(sum < 2 && cell == 1){
		return 0
	}

	return -1
}
