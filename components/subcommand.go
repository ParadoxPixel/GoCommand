package components

type SubCommand struct {
	Name   []string
	Syntaxes []*Syntax
	handler func([]interface{}, int)
}

func NewSubCommand(name... string) *SubCommand {
	return &SubCommand{
		Name:     name,
		Syntaxes: []*Syntax{},
		handler: nil,
	}
}

func (sc *SubCommand) AddSyntax(usage string) *Syntax {
	s := &Syntax{
		Usage:     usage,
		Arguments: []Argument{},
	}
	sc.Syntaxes = append(sc.Syntaxes, s)
	return s
}

func (sc *SubCommand) SetHandler(f func([]interface{}, int)) {
	sc.handler = f
}

func (sc *SubCommand) Handle(args []interface{}, i int) {
	if sc.handler == nil {
		return
	}

	sc.handler(args, i)
}