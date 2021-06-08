package main

import (
	"errors"
	"github.com/ParadoxPixel/GoCommand/components"
	"github.com/ParadoxPixel/GoCommand/objects"
)

type CommandFactory struct {
	Name string
	Middlewares []components.Middleware
	m *objects.CommandMap
}

func NewFactory(name string) *CommandFactory {
	return &CommandFactory{
		Name:        name,
		Middlewares: []components.Middleware{},
		m:           objects.NewCommandMap(),
	}
}

func (cf *CommandFactory) AddMiddleware(middleware components.Middleware) {
	cf.Middlewares = append(cf.Middlewares, middleware)
}

func (cf *CommandFactory) AddCommand(cmd *components.SubCommand) {
	cf.m.Add(cmd)
}

func (cf *CommandFactory) Handle(args []string) error {
	cmd := cf.m.Get(args)
	if cmd == nil {
		return errors.New("no sub-command found")
	}

	args = args[len(cmd.Name):]
	for _,middleware := range cf.Middlewares {
		if !middleware.Continue(cmd) {
			return errors.New("failed to pass middleware")
		}
	}

	var b,ok bool
	str := ""
	var i,j int
	var arguments []components.Argument
	var parsed []interface{}

	for index,syntax := range cmd.Syntaxes {
		arguments = syntax.Arguments
		if len(arguments) > len(args) {
			continue
		}

		b = true
		j = 0
		parsed = []interface{}{}
		for _,argument := range arguments {
			array := args[j:]
			ok,i = argument.Check(array, parsed)
			if !ok {
				str = argument.Message(array)
				b = false
				break
			}

			parsed = append(parsed, argument.Get(array[:i], parsed))
			j += i
		}

		if j != len(args) || !b {
			continue
		}

		cmd.Handle(parsed, index)
		return nil
	}

	if str != "" {
		return errors.New(str)
	}

	return errors.New("usage: "+cmd.Syntaxes[0].Usage)
}