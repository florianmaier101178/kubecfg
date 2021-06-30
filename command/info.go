package command

import (
	"fmt"
	"strings"
)

type InfoCommand struct {
}

func (i *InfoCommand) Help() string {
	helpText := `
Usage: kubecfg info

  General information.
`
	return strings.TrimSpace(helpText)
}

func (i *InfoCommand) Run(args []string) int {
	if len(args) > 0 {
		fmt.Println(i.Help())
		return 1
	}
	//TODO write info about az login, az account set, ...
	//write installation info's where binary is located, about .kubecfg, maybe autocompletion
	return 0
}

func (i *InfoCommand) Synopsis() string {
	return "General information"
}
