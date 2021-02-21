package poker

import "io"

// Game manages the state of a game.
type Game interface {
	Start(numberOfPlayers int, destination io.Writer)
	Finish(winner string)
}
