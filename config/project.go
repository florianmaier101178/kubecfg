package config

import (
	"errors"
	"fmt"
)

type Context string

type Project struct {
	Name            string    `json:"name"`
	SelectedContext Context   `json:"selectedContext"`
	Contexts        []Context `json:"contexts"`
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

func (p *Project) RemoveContext(context Context) (*Project, error) {
	existingContext, i := p.existingContext(context)
	if existingContext {
		p.Contexts = append(p.Contexts[:i], p.Contexts[i+1:]...)
		if p.SelectedContext == context {
			p.UnselectContext()
		}
		return p, nil
	}
	return p, errors.New(
		fmt.Sprintf("given context: '%s' not existing for project: '%s'", context, p.Name))
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

func (p *Project) isContextSelected(context Context) bool {
	if p.SelectedContext == context {
		return true
	}
	return false
}

func (p *Project) SelectContext(context Context) (*Project, error) {
	existingContext, _ := p.existingContext(context)
	if existingContext {
		p.SelectedContext = context
		return p, nil
	}
	return p, errors.New(
		fmt.Sprintf("given context: '%s' not existing for project: '%s'", context, p.Name))
}

func (p *Project) UnselectContext() *Project {
	if p.SelectedContext != unselected {
		p.SelectedContext = unselected
	}
	return p
}

func (p *Project) GetContextNames() []Context {
	return p.Contexts
}
