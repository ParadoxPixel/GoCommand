package objects

import "github.com/ParadoxPixel/GoCommand/components"

type CommandMap struct {
	m map[string]interface{}
}

func NewCommandMap() *CommandMap {
	return &CommandMap{
		m: make(map[string]interface{}),
	}
}

func (cm *CommandMap) Get(args []string) *components.SubCommand {
	return cm.get(0, args...)
}

func (cm *CommandMap) Add(cmd *components.SubCommand) {
	cm.add(cmd.Name, 0, cmd)
}

func (cm *CommandMap) get(i int, args ...string) *components.SubCommand {
	if len(args) <= i {
		return cm.get(0, "")
	}

	obj, ok := cm.m[args[i]]
	if !ok {
		obj, ok = cm.m[""]
		if !ok {
			return nil
		}

		return obj.(*components.SubCommand)
	}

	if len(args) == 1 {
		if cmd, ok := obj.(*CommandMap); ok {
			return cmd.get(0, "")
		}

		return obj.(*components.SubCommand)
	}

	if cmd, ok := obj.(*CommandMap); ok {
		return cmd.get(i + 1, args...)
	}

	return nil
}

func (cm *CommandMap) add(name []string, i int, cmd *components.SubCommand) {
	if len(name) <= i {
		if _,ok := cm.m[""]; ok {
			return
		}

		cm.m[""] = cmd
		return
	}

	var m *CommandMap
	if _,ok := cm.m[name[i]]; ok {
		m = cm.m[name[i]].(*CommandMap)
	} else {
		m = NewCommandMap()
		cm.m[name[i]] = m
	}

	m.add(name, i + 1, cmd)
}