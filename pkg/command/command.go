package command

import (
	"github.com/jessevdk/go-flags"
)

func Parse(opts interface{}, arguments []string) ([]string, error) {
	return flags.ParseArgs(opts, arguments)
}
