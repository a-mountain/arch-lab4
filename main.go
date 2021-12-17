package main

import (
	"arch-lab4/engine"
	"arch-lab4/parser"
	"bufio"
	"os"
)

var inputFile = "commands.txt"

func main() {
	eventLoop := engine.NewEventLoop()
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			command := parser.Parse(commandLine)
			eventLoop.Post(command)
		}
	}
	eventLoop.AwaitFinish()
}
