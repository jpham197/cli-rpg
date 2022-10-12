package monsters

import (
	creature "clirpg/main/creatures"
)

type Monster struct {
	creature.Creature
	evil bool
}
