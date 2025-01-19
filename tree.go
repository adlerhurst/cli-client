package cli

import "github.com/spf13/pflag"

type IndexTree struct {
	set      *pflag.FlagSet
	startIdx int
	endIdx   int

	sub []IndexTree
}

func (it *IndexTree) Parse(args []string) error {
	if err := it.set.Parse(args[it.startIdx : it.endIdx+1]); err != nil {
		return err
	}
	for _, sub := range it.sub {
		if err := sub.Parse(args); err != nil {
			return err
		}
	}
	return nil
}
