package chance

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return r.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return r.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	//
	animals := []string{
		"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog",
	}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}
