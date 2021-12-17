package parser

import (
	"arch-lab4/commands"
	"arch-lab4/engine"
	"strings"
)

func Parse(line string) engine.Command {
	tokens := strings.Split(line, " ")
	if len(tokens) < 2 {
		return &commands.PrintCommand{Msg: "SYNTAX ERROR: Command should have command name and at least one argument"}
	}
	name := tokens[0]

	if name == "print" {
		if len(tokens) != 2 {
			return &commands.PrintCommand{Msg: "SYNTAX ERROR: Command `print` should have 1 arg"}
		}
		message := tokens[1]
		return &commands.PrintCommand{Msg: message}
	}

	if name == "delete" {
		if len(tokens) != 3 {
			return &commands.PrintCommand{Msg: "SYNTAX ERROR: Command `delete` should have 2 args"}
		}
		str := tokens[1]
		symbolToDelete := tokens[2]
		return &commands.DeleteCommand{Str: str, SymbolToDelete: symbolToDelete}
	}

	return &commands.PrintCommand{Msg: "SYNTAX ERROR: wrong command name"}
}
