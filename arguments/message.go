package arguments

import (
	"github.com/ParadoxPixel/GoCommand/components"
	"strings"
)

type Message struct {
	components.Argument
}

func (m Message) Message(args[]string) string {
	return "unable message: "+args[0]
}

func (m Message) Check(args[]string, _ []interface{}) (bool,int) {
	if !strings.HasPrefix(args[0], "\"") {
		return true, 1
	}

	for i,str := range args {
		if strings.HasSuffix(str, "\"") {
			return true, i + 1
		}
	}

	return true, len(args)
}

func (m Message) Get(args []string, _ []interface{}) interface{} {
	return strings.ReplaceAll(strings.Join(args, " "), "\"", "")
}