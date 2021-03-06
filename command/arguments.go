package command

import (
	"errors"
	"kubecfg/config"
)

type ArgumentWithOptionalProjectName interface {
	Name() string
	Available() bool
}

type OptionalProjectNameArgument struct {
	name      string
	available bool
}

func (o OptionalProjectNameArgument) Name() string {
	return o.name
}

func (o OptionalProjectNameArgument) Available() bool {
	return o.available
}

func ParseOptionalProjectNameArgument(args []string) (*OptionalProjectNameArgument, error) {
	optionalProjectNameArg := &OptionalProjectNameArgument{}

	if len(args) > 1 {
		return optionalProjectNameArg, errors.New("invalid number of arguments")
	}

	if len(args) == 1 {
		optionalProjectNameArg.name = args[0]
		optionalProjectNameArg.available = true

		return optionalProjectNameArg, nil
	}

	optionalProjectNameArg.available = false
	return optionalProjectNameArg, nil
}

type ProjectNameAndContextArguments struct {
	OptionalProjectNameArgument
	ContextName          config.Context
}

func (p ProjectNameAndContextArguments) Name() string {
	return p.name
}

func (p ProjectNameAndContextArguments) ProjectName() string {
	return p.Name()
}

func (p ProjectNameAndContextArguments) Available() bool {
	return p.available
}

func ParseProjectNameAndContextArguments(args []string) (*ProjectNameAndContextArguments, error) {
	projectNameAndContextArgs := &ProjectNameAndContextArguments{}

	if len(args) > 3 {
		return projectNameAndContextArgs, errors.New("invalid number of arguments")
	}

	projectNameAndContextArgs.ContextName = config.Context(args[0])

	projectName, err := extractProjectName(args)
	if err != nil {
		projectNameAndContextArgs.available = false
		return projectNameAndContextArgs, nil
	}
	projectNameAndContextArgs.name = projectName
	projectNameAndContextArgs.available = true

	return projectNameAndContextArgs, nil
}

func extractProjectName(args []string) (string, error) {
	for i, value := range args {
		if value == "-project" {
			return args[i+1], nil
		}
	}
	return "", errors.New("cannot parse -project argument")
}

type ContextAddArguments struct {
	ContextName   config.Context
	ResourceGroup string
	ClusterName   string
	ProjectName   string
}

func ParseContextAddArguments(args []string) (*ContextAddArguments, error) {
	contextAddArgs := &ContextAddArguments{}

	if len(args) != 7 {
		return contextAddArgs, errors.New("invalid number of arguments")
	}

	contextAddArgs.ContextName = config.Context(args[0])

	projectName, err := extractProjectName(args)
	if err != nil {
		return contextAddArgs, err
	}
	contextAddArgs.ProjectName = projectName

	rg, err := extractResourceGroup(args)
	if err != nil {
		return contextAddArgs, err
	}
	contextAddArgs.ResourceGroup = rg

	clusterName, err := extractClusterName(args)
	if err != nil {
		return contextAddArgs, err
	}
	contextAddArgs.ClusterName = clusterName

	return contextAddArgs, nil
}

func extractResourceGroup(args []string) (string, error) {
	for i, value := range args {
		if value == "-rg" {
			return args[i+1], nil
		}
	}
	return "", errors.New("cannot parse -rg argument")
}

func extractClusterName(args []string) (string, error) {
	for i, value := range args {
		if value == "-context" {
			return args[i+1], nil
		}
	}
	return "", errors.New("cannot parse -context argument")
}
