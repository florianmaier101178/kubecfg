package config

import "testing"

func TestNewConfig(t *testing.T) {
	config := NewConfig()

	t.Run("verify selected project", func(t *testing.T) {
		expected := "unselected"

		if config.SelectedProject != expected {
			t.Errorf("got %v, but expected %v", config.SelectedProject, expected)
		}
	})

	t.Run("verify empty project slice", func(t *testing.T) {
		expected := 0

		if len(config.Projects) != expected {
			t.Errorf("got %d, but expected %d", len(config.Projects), expected)
		}
	})
}

func configWithAddedProject(projectName string) *Config {
	c := NewConfig()
	c.AddProject(*NewProject(projectName))
	return c
}

func TestAddProject(t *testing.T) {
	c := configWithAddedProject("business")

	t.Run("verify length of project slice", func(t *testing.T) {
		expected := 1

		if len(c.Projects) != expected {
			t.Errorf("got %d, but expected %d", len(c.Projects), expected)
		}
	})

	t.Run("verify existence of added project", func(t *testing.T) {
		projectName := "business"

		context, _ := c.existingProject(projectName)
		if !context {
			t.Errorf("project not existing in contexts of project")
		}
	})
}

func TestAddProjectNotPossibleIfProjectAlreadyExisting(t *testing.T) {
	c := configWithAddedProject("business")
	expected := "project: 'business' already existing in config"

	_, err := c.AddProject(*NewProject("business"))

	if err.Error() != expected {
		t.Errorf("got \"%s\", but expected: \"%s\" ", expected, err.Error())
	}
}

func configWithMultipleProjects() *Config {
	c := configWithAddedProject("business")
	c.AddProject(*NewProject("components"))
	return c
}

func TestAddProjectForMultipleProjectsExistent(t *testing.T) {
	c := configWithMultipleProjects()
	c.AddProject(*NewProject("example"))

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 3

		if len(c.Projects) != expected {
			t.Errorf("got %d, but expected %d", len(c.Projects), expected)
		}
	})

	t.Run("verify existence of added project", func(t *testing.T) {
		projectName := "example"

		project, _ := c.existingProject(projectName)
		if !project {
			t.Errorf("project not existing projects of config")
		}
	})
}
