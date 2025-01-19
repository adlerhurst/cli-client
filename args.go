package cli

type PrimitiveFlagger interface {
	IsPrimitive(flagName string) bool
}

type NestedFlagger interface {
	FlaggerByName(flagName string) Flagger
}

type Flagger interface{
	PrimitiveFlags() []string
	// NestedFlagger
}

type ObjectArgs[T Flagger] struct{
	Object *T
	Args []string
}

func ParseArgs[T any](args []string) *ObjectArgs[T]{
	return nil
}

func ParseArgs(flagger Flagger, args []string) (primitiveArgs, nestedArgs map[string][][]string){
	primitiveArgs = PrimitiveArgs(flagger, args)
	nestedArgs = NestedArgs(flagger, args)
	return primitiveArgs, nestedArgs
}

func PrimitiveArgs(flagger PrimitiveFlagger, args []string) []string {
	for i, arg := range args {
		if arg[0] != '-' {
			continue
		}
		if arg[1] != '-' {
			// TODO: handle shorthand
			continue
		}
		if flagger.IsPrimitive(arg[2:]) {
			continue
		}
		return args[:i]
	}
	return args
}

func NestedArgs(flagger NestedFlagger, args []string) map[string][][]string{
	indexes := make(map[string][][]string)
	var (
		currentFlag string
		currentArgs []string
	)

	for _, arg := range args {
		if arg[0] != '-' {
			currentArgs = append(currentArgs, arg)
			continue
		}
		if arg[1] != '-' {
			// TODO: handle shorthand
			continue
		}
		nestedFlagger := flagger.FlaggerByName(arg[2:])
		// TODO: handle nil

		nestedArgs 
}