package arguments

import "errors"

func extractProjectName(args []string) (string, error) {
	for i, value := range args {
		if value == "-project" {
			return args[i+1], nil
		}
	}
	return "", errors.New("cannot parse -project argument")
}
