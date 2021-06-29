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

		context, _ := c.ExistingProject(projectName)
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

func TestAddProjectForMultipleProjectsExisting(t *testing.T) {
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

		project, _ := c.ExistingProject(projectName)
		if !project {
			t.Errorf("project not existing projects of config")
		}
	})
}

func TestRemoveProject(t *testing.T) {
	c := configWithAddedProject("business")
	c.RemoveProject("business")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 0

		if len(c.Projects) != expected {
			t.Errorf("got %d, but expected %d", len(c.Projects), expected)
		}
	})

	t.Run("verify that context is not existent in contexts", func(t *testing.T) {
		projectName := "business"

		project, _ := c.ExistingProject(projectName)
		if project {
			t.Errorf("project existing in projects of config")
		}
	})
}

func TestRemoveSelectedProject(t *testing.T) {
	c := configWithAddedProject("business")
	c.AddProject(*NewProject("components"))
	c.SelectProject("business")
	c.RemoveProject("business")

	if c.SelectedProject != unselected {
		t.Errorf("project '%v' is the selected project", c.SelectedProject)
	}
}

func TestRemoveUnselectedProject(t *testing.T) {
	c := configWithAddedProject("business")
	c.AddProject(*NewProject("components"))
	c.SelectProject("business")
	c.RemoveProject("components")

	if c.SelectedProject != "business" {
		t.Errorf("project '%v' is the selected project", c.SelectedProject)
	}
}

func TestRemoveProjectNotPossibleIfNoMatchingProjectExists(t *testing.T) {
	c := configWithAddedProject("business")
	c.RemoveProject("components")

	t.Run("verify length of project slice", func(t *testing.T) {
		expected := 1

		if len(c.Projects) != expected {
			t.Errorf("got %d, but expected %d", len(c.Projects), expected)
		}
	})

	t.Run("verify that project is not existent in projects", func(t *testing.T) {
		expected := "business"

		context, _ := c.ExistingProject(expected)
		if !context {
			t.Errorf("project not existing in projects of config")
		}
	})
}

func TestRemoveProjectWithMultipleProjectsExisting(t *testing.T) {
	c := configWithMultipleProjects()
	c.RemoveProject("business")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 1

		if len(c.Projects) != expected {
			t.Errorf("got %d, but expected %d", len(c.Projects), expected)
		}
	})

	t.Run("verify that context is not existent in contexts", func(t *testing.T) {
		projectName := "business"

		context, _ := c.ExistingProject(projectName)
		if context {
			t.Errorf("project existing in projects of config")
		}
	})
}

func configWithSelectedProject(projectName string) *Config {
	c := configWithAddedProject(projectName)
	c.SelectedProject = projectName
	return c
}

func TestIsProjectSelected(t *testing.T) {
	businessSelected := configWithSelectedProject("business")
	businessSelected.AddProject(*NewProject("components"))

	t.Run("project is selected", func(t *testing.T) {
		expected := "business"

		if !businessSelected.IsProjectSelected(expected) {
			t.Errorf("project '%v' is not the selected project", expected)
		}
	})

	t.Run("project is not selected", func(t *testing.T) {
		expected := "components"

		if businessSelected.IsProjectSelected(expected) {
			t.Errorf("project '%v' is the selected project", expected)
		}
	})
}

func TestUnselectProject(t *testing.T) {
	config := configWithSelectedProject("business")
	config.UnselectProject()
	expected := "unselected"

	if config.SelectedProject != expected {
		t.Errorf("got %v, but expected %v", config.SelectedProject, expected)
	}
}

func TestUnselectProjectForAlreadyUnselectedProject(t *testing.T) {
	config := configWithAddedProject("business")
	config.UnselectProject()
	expected := "unselected"

	if config.SelectedProject != expected {
		t.Errorf("got %v, but expected %v", config.SelectedProject, expected)
	}
}

func TestSelectProject(t *testing.T) {
	config := configWithMultipleProjects()
	config.SelectProject("components")

	expected := "components"

	if config.SelectedProject != expected {
		t.Errorf("got %v, but expected %v", config.SelectedProject, expected)
	}
}

func TestSelectedProjectNotPossibleIfProjectNotExisting(t *testing.T) {
	config := configWithMultipleProjects()
	expected := "given project: 'example' not existing in config"

	_, err := config.SelectProject("example")

	if err.Error() != expected {
		t.Errorf("got \"%s\", but expected: \"%s\" ", expected, err.Error())
	}
}
