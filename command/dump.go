package command

import (
	"fmt"
	"strings"
)

type DumpCommand struct {
}

func (d *DumpCommand) Help() string {
	helpText := `
Usage: kubecfg dump

  Dump configuration as tar.gz file.
`
	return strings.TrimSpace(helpText)
}

//TODO implement, maybe links are kind of problematic, think about export and import functionality instead
func (d *DumpCommand) Run(args []string) int {
	if len(args) > 0 {
		fmt.Println(d.Help())
		return 1
	}
	return 0
}

func (d *DumpCommand) Synopsis() string {
	return "Dump configuration as tar.gz file"
}
