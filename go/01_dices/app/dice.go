package app

import "math/rand"

type Dice struct{}

func (Dice) Roll() int { return rand.Intn(5) + 1 }
