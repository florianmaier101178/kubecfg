package config

import (
	"errors"
	"fmt"
)

const unselected = "unselected"

type Config struct {
	SelectedProject string
	Projects        []Project
}

func NewConfig() *Config {
	return &Config{
		SelectedProject: unselected,
		Projects:        []Project{},
	}
}

func (c *Config) AddProject(project Project) (*Config, error) {
	existingProject, _ := c.existingProject(project.Name)
	if !existingProject {
		c.Projects = append(c.Projects, project)
		return c, nil
	}
	return c, errors.New(
		fmt.Sprintf("project: '%s' already existing in config", project.Name))
}

func (c *Config) RemoveProject(projectName string) *Config {
	existingProject, i := c.existingProject(projectName)
	if existingProject {
		c.Projects = append(c.Projects[:i], c.Projects[i+1:]...)
	}
	return c
}

func (c *Config) existingProject(projectName string) (bool, int) {
	if len(c.Projects) > 0 {
		for i, p := range c.Projects {
			if p.Name == projectName {
				return true, i
			}
		}
	}
	return false, -1
}
