package game_of_life_5

import "errors"

func IsAlive(cell int, sum int)int {

	if sum < 0 {
		panic(errors.New("Invalid neighbours array"))
	}

	if cell == 0 {
		if sum == 3 {
			return 1
		}

		return 0
	}
	if sum < 2 {
		return 0
	}

	if sum <= 3 {
		return 1
	}

	return 0
}
