package arguments

import (
	"errors"
)

type ContextUnselectArguments struct {
	ProjectNameArgument
}

func ParseContextUnselectArguments(args []string) (*ContextUnselectArguments, error) {
	contextUnselectArgs := &ContextUnselectArguments{}

	if len(args) > 1 {
		return contextUnselectArgs, errors.New("invalid number of arguments")
	}

	if len(args) == 1 {
		contextUnselectArgs.ProjectName = args[0]
		contextUnselectArgs.ProjectNameAvailable = true

		return contextUnselectArgs, nil
	}

	contextUnselectArgs.ProjectNameAvailable = false
	return contextUnselectArgs, nil
}

func (c ContextUnselectArguments) Name() string {
	return c.ProjectName
}

func (c ContextUnselectArguments) Available() bool {
	return c.ProjectNameAvailable
}
