package arguments

import (
	"errors"
	"kubecfg/config"
)

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
