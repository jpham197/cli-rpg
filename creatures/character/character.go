package character

import (
	creature "clirpg/main/creatures"
	"math/rand"
)

// Main Character
type Character struct {
	creature.Creature
}

func (c *Character) move() int64 {
	// generates random event
	eventGenerator := rand.Intn(100)
	return int64(eventGenerator)
}
