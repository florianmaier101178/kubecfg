package config

import (
	"errors"
	"fmt"
)

type Context string

type Project struct {
	Name            string
	SelectedContext Context
	Contexts        []Context
}

func NewProject(name string) *Project {
	return &Project{
		Name:            name,
		SelectedContext: unselected,
		Contexts:        []Context{},
	}
}

func (p *Project) AddContext(context Context) (*Project, error) {
	existingContext, _ := p.existingContext(context)
	if !existingContext {
		p.Contexts = append(p.Contexts, context)
		return p, nil
	}
	return p, errors.New(
		fmt.Sprintf("context: '%s' already existing for project: '%s'", context, p.Name))
}

func (p *Project) RemoveContext(context Context) *Project {
	existingContext, i := p.existingContext(context)
	if existingContext {
		p.Contexts = append(p.Contexts[:i], p.Contexts[i+1:]...)
	}
	return p
}

func (p *Project) existingContext(context Context) (bool, int) {
	if len(p.Contexts) > 0 {
		for i, c := range p.Contexts {
			if c == context {
				return true, i
			}
		}
	}
	return false, -1
}
