package command

import (
	"testing"
)

type opts struct {
	Name string `short:"n" long:"name" description:"姓名"`
}

func Test_Parse(t *testing.T) {
	var o = opts{}
	args, _ := Parse("doTest", &o, []string{"doSomething", "-n", "test"})
	println(len(args))
	println(o.Name)
}
