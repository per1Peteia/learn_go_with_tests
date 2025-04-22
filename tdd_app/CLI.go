package poker

import ()

type CLI struct {
	store PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.store.RecordWin("Cleo")
}
