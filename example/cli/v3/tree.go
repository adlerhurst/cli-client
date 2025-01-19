package cli

import "github.com/spf13/pflag"

type indexTree struct {
	flags    *pflag.FlagSet
	startIdx int
	endIdx   int

	sub []indexTree
}

func (it *indexTree) parse(args []string) error {
	if err := it.flags.Parse(args[it.startIdx : it.endIdx+1]); err != nil {
		return err
	}
	for _, sub := range it.sub {
		if err := sub.parse(args); err != nil {
			return err
		}
	}
	return nil
}
