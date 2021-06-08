package arguments

import (
	"github.com/ParadoxPixel/GoCommand/components"
)

type String struct {
	Valid []string
	components.Argument
}

func (s String) Message(args[]string) string {
	return "unable string value: "+args[0]
}

func (s String) Check(args[]string, _ []interface{}) (bool,int) {
	if s.Valid == nil || len(s.Valid) == 0 {
		return true, 1
	}

	for _,str := range s.Valid {
		if str == args[0] {
			return true, 1
		}
	}

	return false, 1
}

func (s String) Get(args []string, _ []interface{}) interface{} {
	return args[0]
}