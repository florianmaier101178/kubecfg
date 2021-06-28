package command

import (
	"fmt"
	"strings"
)

type VersionCommand struct {
	Version string
}

func (c *VersionCommand) Help() string {
	helpText := `
Usage: kubecfg version

  Displays the version.
`
	return strings.TrimSpace(helpText)
}

func (c *VersionCommand) Run(args []string) int {
	fmt.Printf("kubecfg version: %s\n", c.Version)
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Show the current version"
}
