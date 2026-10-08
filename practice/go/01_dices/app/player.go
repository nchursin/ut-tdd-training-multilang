package app

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
