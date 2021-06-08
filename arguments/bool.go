package arguments

import (
	"github.com/ParadoxPixel/GoCommand/components"
	"strconv"
)

type Bool struct {
	components.Argument
}

func (b Bool) Message(args[]string) string {
	return "unable to parse \""+args[0]+"\" to bool"
}

func (b Bool) Check(args[]string, _ []interface{}) (bool,int) {
	_,ok := strconv.ParseBool(args[0])
	return ok == nil, 1
}

func (b Bool) Get(args []string, _ []interface{}) interface{} {
	j,_ := strconv.ParseBool(args[0])
	return j
}