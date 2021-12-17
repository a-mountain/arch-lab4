package commands

import (
	"arch-lab4/engine"
	"strings"
)

type DeleteCommand struct {
	Str            string
	SymbolToDelete string
}

func (c *DeleteCommand) Execute(handler engine.Handler) {
	res := strings.ReplaceAll(c.Str, c.SymbolToDelete, "")
	handler.Post(&PrintCommand{Msg: res})
}
