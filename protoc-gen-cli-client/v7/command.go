package v7

type command interface {
	Execute() error
}

type batch struct {
	cmds []command
}

func (b batch) Execute() error {
	for _, cmd := range b.cmds {
		if err := cmd.Execute(); err != nil {
			return err
		}
	}
	return nil
}
