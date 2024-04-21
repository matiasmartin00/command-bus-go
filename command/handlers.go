package command

import "log"

type CommandHandler interface {
	Execute(Command)
}

type OneCommandHandler struct {
}

func (o *OneCommandHandler) Execute(c Command) {
	cmd, ok := c.(*OneCommand)
	if !ok {
		log.Fatalf("can't handle command %s\n", c.GetID())
		return
	}
	log.Printf("handling command one -> %v\n", cmd)
}

type TwoCommandHandler struct {
}

func (o *TwoCommandHandler) Execute(c Command) {
	cmd, ok := c.(*TwoCommand)
	if !ok {
		log.Fatalf("can't handle command %s\n", c.GetID())
		return
	}
	log.Printf("handling command two -> %v\n", cmd)
}

type ThreeCommandHandler struct {
}

func (o *ThreeCommandHandler) Execute(c Command) {
	cmd, ok := c.(*ThreeCommand)
	if !ok {
		log.Fatalf("can't handle command %s\n", c.GetID())
		return
	}
	log.Printf("handling command three -> %v\n", cmd)
}
