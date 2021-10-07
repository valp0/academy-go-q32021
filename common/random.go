package common

import (
	"math/rand"
	"time"
)

// Returns a positive number between 0 and the given integer argument, safe for pokéAPI pokémon ID.
func RandInt(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if num == 0 || num == 899 {
		return RandInt(max)
	}

	return num
}
