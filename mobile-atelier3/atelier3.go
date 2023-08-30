package atelier3

import "math/rand"

const MAX_INT = 2147483647

func ArrayGenerate(lenght int) []int {
	var array [MAX_INT]int
	var slice []int = array[0 : lenght+1]
	var random int

	for i := 0; i < lenght-1; i++ {
		random = rand.Intn(10000)
		random++
		slice[i] = random
	}

	return slice
}
