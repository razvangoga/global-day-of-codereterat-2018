package game_of_life_2

func IsAlive(cell int, neighbours [8]int) int {
	
	sum := 0

	for i := 0; i < len(neighbours); i++ {
		sum += neighbours[i]
	}

	if cell == 0 {
		if sum == 3 {
			return 1
		} else {
			return 0
		}
	}

	if sum < 2 {
		return 0
	}
	if sum >=2 && sum <=3 {
		return 1
	}
	if sum >3 {
		return 0
	}

	return -1
}
