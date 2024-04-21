package command

import (
	"log"
	"sync"
)

var instance commandBus
var once sync.Once

func init() {
	once.Do(func() {
		instance = commandBus{
			handlers: map[string]CommandHandler{
				OneCmdID:   &OneCommandHandler{},
				TwoCmdID:   &TwoCommandHandler{},
				ThreeCmdID: &ThreeCommandHandler{},
			},
		}
	})
}

func GetCommandBus() commandBus {
	return instance
}

type commandBus struct {
	handlers map[string]CommandHandler
}

func (c *commandBus) Exec(cmd Command) {
	h, ok := c.handlers[cmd.GetID()]
	if !ok {
		log.Fatalf("without handler for command %s\n", cmd.GetID())
		return
	}

	h.Execute(cmd)
}
