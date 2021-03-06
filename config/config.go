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
		return c, nil
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

func (c *Config) HasSelectedProject() bool {
	if c.SelectedProject != unselected {
		return true
	}
	return false
}

func (c *Config) IsProjectSelected(projectName string) bool {
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

func (c *Config) GetProject(projectName string) (*Project, error) {
	existingProject, _ := c.ExistingProject(projectName)
	if existingProject {
		for _, project := range c.Projects {
			if project.Name == projectName {
				return &project, nil
			}
		}
	}
	return &Project{}, errors.New(fmt.Sprintf("No project found for projectName: '%s'\n", projectName))
}

func (c *Config) GetSelectedProject() (*Project, error) {
	err := errors.New("There is no selected project available")

	if !c.HasSelectedProject() {
		return nil, err
	}
	for _, project := range c.Projects {
		if c.SelectedProject == project.Name {
			return &project, nil
		}
	}
	return nil, err
}

func (c *Config) UpdateProject(updatedProject Project) (*Config, error) {
	existingProject, _ := c.ExistingProject(updatedProject.Name)
	if existingProject {
		for i, project := range c.Projects {
			if project.Name == updatedProject.Name {
				c.Projects[i] = updatedProject
				return c, nil
			}
		}
	}
	return c, errors.New(
		fmt.Sprintf("given project: '%s' not existing in config", updatedProject.Name))
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
