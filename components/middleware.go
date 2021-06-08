package components

type Middleware interface {

	Continue(*SubCommand) bool

}