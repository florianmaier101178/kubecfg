# kubecfg
A simple management tool to handle kubernetes configurations for multiple projects with different contexts. Right now only Azure AKS is supported as kubernetes cloud provider.

This tool got developed for **personal usage**.

There will be no support for it. I will not guarantee that it stays backwards compatible.

## Installation
``` bash
$ make install
-> /usr/local/bin/kubecfg
```

## Usage
``` bash
$ kubecfg
Usage: kubecfg [--version] [--help] <command> [<args>]

Available commands are:
    context    Context management
    info       General information
    init       Initialize kubecfg tool
    project    Project management
    version    Show the current version
```

## Commands

### info
``` bash
$ kubecfg info -h
Usage: kubecfg info

  General information.
```

### version
``` bash
$ kubecfg version -h
Usage: kubecfg version

  Displays the version.
```

### init
``` bash
$ kubecfg init --help
Usage: kubecfg init

  Initializes the kubecfg tool. It creates a ~/.kubecfg directory
  to store all configuration needed for kubecfg.
```

### project
``` bash
$ kubecfg project --help
Usage: kubecfg project

  Add, delete, list, select and show projects.

Subcommands:
    add         Add a new project
    list        List projects
    remove      Remove a project
    select      Select a project
    show        Show selected project
    unselect    Unselect project choice
```

### context
``` bash
$ kubecfg context -h
Usage: kubecfg context

  Add, delete, list, select and show contexts for the currently selected project.

Subcommands:
    add         Add a new context for the currently selected project
    list        List contexts
    remove      Remove a context from a project
    select      Select a context for a project
    show        Show selected context of selected project
    unselect    Unselect context choice for a project
```

## Autocompletion
* I know that the used lib for cli supports autocompletion. 
* Due to time constraints I decided to use a simple bash_completion script to autocomplete.

## Internals
* The kubectl command normally stores its contexts inside the ~/.kube/config file.
* *kubecfg* uses its own configuration directory ~/.kubecfg.
* Inside ~/.kubecfg is a config file / symlink to project specific kubernetes config files.
* The original ~/.kube/config file has to be replaced with a symlink to ~/.kubecfg/config.
* A setup description is available with the *kubecfg info* command.
