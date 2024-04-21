package command

const (
	OneCmdID   = "one-command-id"
	TwoCmdID   = "two-command-id"
	ThreeCmdID = "three-command-id"
)

type Command interface {
	GetID() string
}

type OneCommand struct {
	Name  string
	Color string
}

func (o *OneCommand) GetID() string {
	return OneCmdID
}

type TwoCommand struct {
	Name  string
	Color string
}

func (o *TwoCommand) GetID() string {
	return TwoCmdID
}

type ThreeCommand struct {
	Name  string
	Color string
}

func (o *ThreeCommand) GetID() string {
	return ThreeCmdID
}
