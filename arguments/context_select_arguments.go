package arguments

import (
	"errors"
	"kubecfg/config"
)

type ContextSelectArguments struct {
	ProjectNameAndContextArguments
}

func ParseContextSelectArguments(args []string) (*ContextSelectArguments, error) {
	contextSelectArgs := &ContextSelectArguments{}

	if len(args) > 3 {
		return contextSelectArgs, errors.New("invalid number of arguments")
	}

	contextSelectArgs.ContextName = config.Context(args[0])

	projectName, err := extractProjectName(args)
	if err != nil {
		contextSelectArgs.ProjectNameAvailable = false
		return contextSelectArgs, nil
	}
	contextSelectArgs.ProjectName = projectName
	contextSelectArgs.ProjectNameAvailable = true

	return contextSelectArgs, nil
}

func (c ContextSelectArguments) Name() string {
	return c.ProjectName
}

func (c ContextSelectArguments) Available() bool {
	return c.ProjectNameAvailable
}
