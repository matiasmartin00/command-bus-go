package main

import "github.com/matiasmartin00/command-bus-go/command"

func main() {
	cb := command.GetCommandBus()

	cmdOne := &command.OneCommand{
		Name:  "Carl",
		Color: "Blue",
	}

	cmdTwo := &command.TwoCommand{
		Name:  "Peter",
		Color: "White",
	}

	cmdThree := &command.ThreeCommand{
		Name:  "Jon",
		Color: "Red",
	}

	cb.Exec(cmdThree)
	cb.Exec(cmdOne)
	cb.Exec(cmdTwo)
}
