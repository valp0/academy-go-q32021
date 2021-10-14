package common

import (
	"math/rand"
	"time"
)

const maxId = 898
const minId = 1

// Returns a positive number between 0 and the given integer argument.
var RandInt = func(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if minId > num || num > maxId {
		return randInt(max)
	}

	return num
}

func randInt(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if minId > num || num > maxId {
		return randInt(max)
	}

	return num
}
