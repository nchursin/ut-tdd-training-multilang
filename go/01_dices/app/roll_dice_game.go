package app

type RollDiceGame struct {
	playersCount int
	bets         []struct {
		player *Player
		chips  Chip
		score  int
	}
}

func (g *RollDiceGame) AddPlayer() {
	if g.playersCount == 6 {
		panic("too many players")
	}
	g.playersCount++
}
func (g *RollDiceGame) RemovePlayer() { g.playersCount-- }
func (g *RollDiceGame) Bet(player *Player, bet Bet) {
	if !player.Has(bet.Chips) {
		panic("invalid operation")
	}
	g.bets = append(g.bets, struct {
		player *Player
		chips  Chip
		score  int
	}{player, bet.Chips, bet.Score})
	player.Take(bet.Chips)
}
func (g *RollDiceGame) Play() {
	winningScore := Dice{}.Roll()
	for _, bet := range g.bets {
		if bet.score == winningScore {
			bet.player.Win(bet.chips.Multiply(6))
		}
	}
}
