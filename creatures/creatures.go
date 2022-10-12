package creatures

import (
	stats "clirpg/main/creatures/stats"
)

type Creature struct {
	stats.Stats
	name string
}

/*
Returns 2 values
bool: Whether the target is dead after the attack
int64: Target's remaining health after being attacked
*/
func (c *Creature) attack(target *Creature) (bool, int64) {
	remainingHealth := target.Health + c.Atk - (target.Def / 2)
	return remainingHealth <= 0, remainingHealth
}
