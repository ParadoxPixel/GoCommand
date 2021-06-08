package arguments

import (
	"github.com/ParadoxPixel/GoCommand/components"
	"strconv"
)

type Int struct {
	Valid []int
	components.Argument
}

func (i Int) Message(args[]string) string {
	return "unable to parse \""+args[0]+"\" to int"
}

func (i Int) Check(args[]string, _ []interface{}) (bool,int) {
	j,ok := strconv.Atoi(args[0])
	if ok != nil {
		return false, 1
	}

	if i.Valid == nil || len(i.Valid) == 0 {
		return true, 1
	}

	for _,str := range i.Valid {
		if str == j {
			return true, 1
		}
	}

	return false, 1
}

func (i Int) Get(args []string, _ []interface{}) interface{} {
	j,_ := strconv.Atoi(args[0])
	return j
}