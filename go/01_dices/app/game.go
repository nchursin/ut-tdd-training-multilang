package app

import "math/rand"

// Game is the minimal game contract used by Player.
type Game interface {
	AddPlayer()
	RemovePlayer()
}

type Bet struct {
	Chips Chip
	Score int
}

type Chip struct{ Amount int }

func (c Chip) Equal(other Chip) bool          { return c.Amount == other.Amount }
func (c Chip) GreaterOrEqual(other Chip) bool { return c.Amount >= other.Amount }
func (c Chip) LessOrEqual(other Chip) bool    { return c.Amount <= other.Amount }
func (c Chip) Add(other Chip) Chip            { return Chip{Amount: c.Amount + other.Amount} }
func (c Chip) Subtract(other Chip) Chip       { return Chip{Amount: c.Amount - other.Amount} }
func (c Chip) Multiply(value int) Chip        { return Chip{Amount: c.Amount * value} }

type Dice struct{}

func (Dice) Roll() int { return rand.Intn(5) + 1 }

type Player struct {
	currentGame    Game
	availableChips Chip
}

func (p *Player) IsInGame() bool { return p.currentGame != nil }
func (p *Player) Join(game Game) {
	if p.IsInGame() {
		panic("invalid operation")
	}
	p.currentGame = game
	game.AddPlayer()
}
func (p *Player) LeaveGame() {
	if !p.IsInGame() {
		panic("invalid operation")
	}
	p.currentGame.RemovePlayer()
	p.currentGame = nil
}
func (p *Player) Buy(chips Chip)      { p.availableChips = chips }
func (p *Player) Has(chips Chip) bool { return p.availableChips.GreaterOrEqual(chips) }
func (p *Player) Take(chips Chip)     { p.availableChips = p.availableChips.Subtract(chips) }
func (p *Player) Win(chips Chip)      { p.availableChips = p.availableChips.Add(chips) }

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
