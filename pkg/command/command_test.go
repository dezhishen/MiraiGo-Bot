package command

import (
	"testing"
)

type opts struct {
	Name string `short:"n" long:"name" description:"姓名"`
}

func Test_Parse(t *testing.T) {
	var o = opts{}
	args, _ := Parse(&o, []string{"doSomething", "-n", "test"})
	print(len(args))
	print(o.Name)
}
