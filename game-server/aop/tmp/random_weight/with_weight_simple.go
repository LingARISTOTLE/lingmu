package random_weight

import "math/rand"

func getWeightChoice(weight [][]uint32) uint32 {
	total := uint32(0)
	winner := 0

	for i, v := range weight {
		total += v[1]
		if uint32(rand.Float32()*float32(total)) < v[1] {
			winner = i
		}
	}

	return weight[winner][0]
}
