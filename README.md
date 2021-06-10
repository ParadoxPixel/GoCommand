# GoCommand
Command Handler for Golang with a few basic argument types

Example:
```golang
import (
  goc "github.com/ParadoxPixel/GoCommand"
  gocc "github.com/ParadoxPixel/GoCommand/components"
  "github.com/ParadoxPixel/GoCommand/arguments"
)

//Register base of command
factory := goc.NewFactory("base")

//Create a new sub-command and set it's handler
cmd := gocc.NewSubCommand("arg1")
cmd.SetHandler(func(args []interface{}, i int) {
  fmt.Println("Args:", args, "Syntax:", i)
})

//Add a syntax with it's argument parsers
syntax := cmd.AddSyntax("base arg1 <string>")
syntax.AddArgument(arguments.String{})

//Register sub-command
factory.AddCommand(cmd)

//TODO Register factory's and fetch them based on the first argument("base" in this case)

//Pass arguments after base to the CommandFactory
err := factory.Handle([]string{"arg1", "test"})
if err != nil {
  fmt.Println("Error while handling command:", err)
}
```
