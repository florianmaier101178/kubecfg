package arguments

import (
	"errors"
)

type ContextShowArguments struct {
	ProjectNameArgument
}

func ParseContextShowArguments(args []string) (*ContextShowArguments, error) {
	contextShowArgs := &ContextShowArguments{}

	if len(args) > 1 {
		return contextShowArgs, errors.New("invalid number of arguments")
	}

	if len(args) == 1 {
		contextShowArgs.ProjectName = args[0]
		contextShowArgs.ProjectNameAvailable = true

		return contextShowArgs, nil
	}

	contextShowArgs.ProjectNameAvailable = false
	return contextShowArgs, nil
}

func (c ContextShowArguments) Name() string {
	return c.ProjectName
}

func (c ContextShowArguments) Available() bool {
	return c.ProjectNameAvailable
}
