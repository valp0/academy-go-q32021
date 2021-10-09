package common

import (
	"math/rand"
	"time"
)

// Returns a positive number between 0 and the given integer argument.
var RandInt = func(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if 1 > num || num > 898 {
		return randInt(max)
	}

	return num
}

func randInt(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if 1 > num || num > 898 {
		return randInt(max)
	}

	return num
}
