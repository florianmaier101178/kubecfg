package arguments

import (
	"errors"
	"kubecfg/config"
)

type ProjectNameArgument struct {
	ProjectName          string
	ProjectNameAvailable bool
}

type ProjectNameAndContextArguments struct {
	ProjectNameArgument
	ContextName          config.Context
}

func extractProjectName(args []string) (string, error) {
	for i, value := range args {
		if value == "-project" {
			return args[i+1], nil
		}
	}
	return "", errors.New("cannot parse -project argument")
}
