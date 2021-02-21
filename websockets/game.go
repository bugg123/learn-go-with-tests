package poker

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}
