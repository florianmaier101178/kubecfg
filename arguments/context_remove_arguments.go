package arguments

import (
	"errors"
	"kubecfg/config"
)

type ContextRemoveArguments struct {
	ContextName          config.Context
	ProjectName          string
	ProjectNameAvailable bool
}

func ParseContextRemoveArguments(args []string) (*ContextRemoveArguments, error) {
	contextRemoveArgs := &ContextRemoveArguments{}

	if len(args) > 3 {
		return contextRemoveArgs, errors.New("invalid number of arguments")
	}

	contextRemoveArgs.ContextName = config.Context(args[0])

	projectName, err := extractProjectName(args)
	if err != nil {
		contextRemoveArgs.ProjectNameAvailable = false
		return contextRemoveArgs, nil
	}
	contextRemoveArgs.ProjectName = projectName
	contextRemoveArgs.ProjectNameAvailable = true

	return contextRemoveArgs, nil
}

func (c ContextRemoveArguments) Name() string {
	return c.ProjectName
}

func (c ContextRemoveArguments) Available() bool {
	return c.ProjectNameAvailable
}
