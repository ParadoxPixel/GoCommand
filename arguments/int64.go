package arguments

import (
	"github.com/ParadoxPixel/GoCommand/components"
	"strconv"
)

type Int64 struct {
	components.Argument
}

func (i Int64) Message(args[]string) string {
	return "unable to parse \""+args[0]+"\" to int64"
}

func (i Int64) Check(args[]string, _ []interface{}) (bool,int) {
	_,ok := strconv.ParseInt(args[0], 10, 64)
	return ok == nil, 1
}

func (i Int64) Get(args []string, _ []interface{}) interface{} {
	j,_ := strconv.ParseInt(args[0], 10, 64)
	return j
}