package arguments

import (
	"errors"
)

type ContextListArguments struct {
	ProjectName          string
	ProjectNameAvailable bool
}

func ParseContextListArguments(args []string) (*ContextListArguments, error) {
	contextListArgs := &ContextListArguments{}

	if len(args) > 1 {
		return contextListArgs, errors.New("invalid number of arguments")
	}

	contextListArgs.ProjectName = args[0]
	contextListArgs.ProjectNameAvailable = true

	return contextListArgs, nil
}

func (c ContextListArguments) Name() string {
	return c.ProjectName
}

func (c ContextListArguments) Available() bool {
	return c.ProjectNameAvailable
}
