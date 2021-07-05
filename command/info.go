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
	fmt.Printf("some information to properly use and setup 'kubecfg'\n\n")
	fmt.Println("1. backup your .kube directory and verify that there is not .kubecfg directory in your home folder")
	fmt.Println("   $ mv ~/.kube ~/.kube_backup")
	fmt.Println("   $ ls -ali ~/.kubecfg  # this command should list nothing")
	fmt.Println("2. login into your azure account")
	fmt.Println("   $ az login")
	fmt.Println("3. initialise kubecfg (should be installed and available at e.g. /usr/local/bin/kubecfg)")
	fmt.Println("   $ kubecfg init")
	fmt.Println("4. list available azure subscriptions")
	fmt.Println("   $ az account show")
	fmt.Println("5. select appropriate azure subscription where AKS to connect to is hosted")
	fmt.Println("   $ az account set --subscription \"HAU Business Dev\"")
	fmt.Println("   $ az account set --subscription \"HAU Digitalisierung Prod\"")
	fmt.Println("   $ az account set --subscription \"LDC Components Dev\"")
	fmt.Println("6. create .kube directory in home folder")
	fmt.Println("   $ mkdir ~/.kube")
	fmt.Println("7. create symbolic link ./kube/config to ./kubecfg/config")
	fmt.Println("   $ ln -s ~/.kubecfg/config ~/.kube/config")
	fmt.Println("")
	fmt.Println("from here on use kubecfg ... have fun!")
	fmt.Println("")
	return 0
}

func (i *InfoCommand) Synopsis() string {
	return "General information"
}
