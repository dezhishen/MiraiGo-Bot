package command

import (
	"github.com/jessevdk/go-flags"
)

func Parse(appName string, opts interface{}, arguments []string) ([]string, error) {
	f := flags.NewParser(opts, flags.Default)
	f.Command.Name = appName
	return f.ParseArgs(arguments)
}
