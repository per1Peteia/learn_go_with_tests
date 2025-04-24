package poker

import ()

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
