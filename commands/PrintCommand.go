package commands

import (
	"arch-lab4/engine"
	"fmt"
)

type PrintCommand struct {
	Msg string
}

func (c *PrintCommand) Execute(handler engine.Handler) {
	fmt.Println(c.Msg)
}
