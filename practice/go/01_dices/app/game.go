package app

// Game is the minimal game contract used by Player.
type Game interface {
	AddPlayer()
	RemovePlayer()
}
