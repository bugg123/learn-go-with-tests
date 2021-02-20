package poker

type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}
