package components

type Argument interface {

	Message([]string) string

	Check([]string, []interface{}) (bool,int)

	Get([]string, []interface{}) interface{}

}
