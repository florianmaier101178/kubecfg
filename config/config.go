package config

import (
	"errors"
	"fmt"
)

const unselected = "unselected"

type Config struct {
	SelectedProject string    `json:"selectedProject"`
	Projects        []Project `json:"projects"`
}

func NewConfig() *Config {
	return &Config{
		SelectedProject: unselected,
		Projects:        []Project{},
	}
}

func (c *Config) AddProject(project Project) (*Config, error) {
	existingProject, _ := c.ExistingProject(project.Name)
	if !existingProject {
		c.Projects = append(c.Projects, project)
		return c, nil
	}
	return c, errors.New(
		fmt.Sprintf("project: '%s' already existing in config", project.Name))
}

func (c *Config) RemoveProject(projectName string) (*Config, error) {
	existingProject, i := c.ExistingProject(projectName)
	if existingProject {
		c.Projects = append(c.Projects[:i], c.Projects[i+1:]...)
		if c.SelectedProject == projectName {
			c.UnselectProject()
		}
	}
	return c, errors.New(
		fmt.Sprintf("given project: '%s' not existing in config", projectName))
}

func (c *Config) ExistingProject(projectName string) (bool, int) {
	if len(c.Projects) > 0 {
		for i, p := range c.Projects {
			if p.Name == projectName {
				return true, i
			}
		}
	}
	return false, -1
}

func (c *Config) isProjectSelected(projectName string) bool {
	if c.SelectedProject == projectName {
		return true
	}
	return false
}

func (c *Config) SelectProject(projectName string) (*Config, error) {
	existingProject, _ := c.ExistingProject(projectName)
	if existingProject {
		c.SelectedProject = projectName
		return c, nil
	}
	return c, errors.New(
		fmt.Sprintf("given project: '%s' not existing in config", projectName))
}

func (c *Config) UnselectProject() *Config {
	if c.SelectedProject != unselected {
		c.SelectedProject = unselected
	}
	return c
}

func (c *Config) GetProjectNames() []string {
	if len(c.Projects) > 0 {
		projectNames := make([]string, 0)
		for _, project := range c.Projects {
			projectNames = append(projectNames, project.Name)
		}
		return projectNames
	}
	return make([]string, 0)
}
