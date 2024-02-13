package cli

import "strings"

type indexes []*index

func (idxs indexes) ByName(names ...string) (res []*index) {
	for _, idx := range idxs {
		for _, name := range names {
			if idx.Flag != name {
				continue
			}
			res = append(res, idx)
			break
		}
	}
	return res
}

func (idxs indexes) LastByName(names ...string) *index {
	for i := len(idxs) - 1; i >= 0; i-- {
		for _, name := range names {
			if idxs[i].Flag != name {
				continue
			}
			return idxs[i]
		}
	}
	return nil
}

func (idxs indexes) Last() *index {
	return idxs[len(idxs)-1]
}

func (idxs indexes) Primitives() *index {
	return idxs.LastByName("")
}

type index struct {
	Flag string
	Args []string
}

func FieldIndexes(args []string, names ...string) (res indexes) {
	// index for primitive fields
	res = append(res, new(index))
	for _, arg := range args {
		flag, found := strings.CutPrefix(arg, "-")
		if !found {
			res[len(res)-1].Args = append(res[len(res)-1].Args, arg)
			continue
		}
		// also trim second -
		flag = strings.TrimPrefix(flag, "-")

		if isFieldFlag(flag, names) {
			res = append(res, &index{Flag: flag})
		} else {
			res[len(res)-1].Args = append(res[len(res)-1].Args, arg)
		}
	}

	return res
}

func isFieldFlag(arg string, names []string) bool {
	for _, name := range names {
		if arg != name {
			continue
		}
		return true
	}
	return false
}
