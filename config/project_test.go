package config

import (
	"testing"
)

func TestNewProject(t *testing.T) {
	project := NewProject("business")

	t.Run("verify project name", func(t *testing.T) {
		expected := "business"

		if project.Name != expected {
			t.Errorf("got %v, but expected %v", project.Name, expected)
		}
	})

	t.Run("verify selected context", func(t *testing.T) {
		var expected Context = "unselected"

		if project.SelectedContext != expected {
			t.Errorf("got %v, but expected %v", project.SelectedContext, expected)
		}
	})

	t.Run("verify empty context slice", func(t *testing.T) {
		expected := 0

		if len(project.Contexts) != expected {
			t.Errorf("got %d, but expected %d", len(project.Contexts), expected)
		}
	})
}

func projectWithAddedContext(context Context) *Project {
	p := NewProject("business")
	p.AddContext(context)
	return p
}

func TestAddContext(t *testing.T) {
	p := projectWithAddedContext("dev")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 1

		if len(p.Contexts) != expected {
			t.Errorf("got %d, but expected %d", len(p.Contexts), expected)
		}
	})

	t.Run("verify existence of added context", func(t *testing.T) {
		var expected Context = "dev"

		context, _ := p.existingContext(expected)
		if !context {
			t.Errorf("context not existing in contexts of project")
		}
	})
}

func TestAddContextNotPossibleIfContextAlreadyExisting(t *testing.T) {
	p := projectWithAddedContext("dev")
	expected := "context: 'dev' already existing for project: 'business'"

	_, err := p.AddContext("dev")

	if err.Error() != expected {
		t.Errorf("got \"%s\", but expected: \"%s\" ", expected, err.Error())
	}
}

func projectWithMultipleContexts() *Project {
	p := projectWithAddedContext("dev")
	p.AddContext("int")
	p.AddContext("stable")
	return p
}

func TestAddContextForMultipleContextsExisting(t *testing.T) {
	p := projectWithMultipleContexts()
	p.AddContext("prod")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 4

		if len(p.Contexts) != expected {
			t.Errorf("got %d, but expected %d", len(p.Contexts), expected)
		}
	})

	t.Run("verify existence of added context", func(t *testing.T) {
		var expected Context = "prod"

		context, _ := p.existingContext(expected)
		if !context {
			t.Errorf("context not existing in contexts of project")
		}
	})
}

func TestRemoveContext(t *testing.T) {
	p := projectWithAddedContext("dev")
	p.RemoveContext("dev")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 0

		if len(p.Contexts) != expected {
			t.Errorf("got %d, but expected %d", len(p.Contexts), expected)
		}
	})

	t.Run("verify that context is not existent in contexts", func(t *testing.T) {
		var expected Context = "dev"

		context, _ := p.existingContext(expected)
		if context {
			t.Errorf("context existing in contexts of project")
		}
	})
}

func TestRemoveSelectedContext(t *testing.T) {
	p := projectWithAddedContext("dev")
	p.AddContext("int")
	p.SelectContext("dev")
	p.RemoveContext("dev")

	if p.SelectedContext != unselected {
		t.Errorf("context '%v' is the selected context", p.SelectedContext)
	}
}

func TestRemoveUnselectedContext(t *testing.T) {
	p := projectWithAddedContext("dev")
	p.AddContext("int")
	p.SelectContext("dev")
	p.RemoveContext("int")

	if p.SelectedContext != "dev" {
		t.Errorf("context '%v' is the selected context", p.SelectedContext)
	}
}

func TestRemoveContextNotPossibleIfNoMatchingContextExists(t *testing.T) {
	p := projectWithAddedContext("dev")
	p.RemoveContext("int")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 1

		if len(p.Contexts) != expected {
			t.Errorf("got %d, but expected %d", len(p.Contexts), expected)
		}
	})

	t.Run("verify that context is not existent in contexts", func(t *testing.T) {
		var expected Context = "dev"

		context, _ := p.existingContext(expected)
		if !context {
			t.Errorf("context not existing in contexts of project")
		}
	})
}

func TestRemoveContextWithMultipleContextsExisting(t *testing.T) {
	p := projectWithMultipleContexts()
	p.RemoveContext("stable")

	t.Run("verify length of context slice", func(t *testing.T) {
		expected := 2

		if len(p.Contexts) != expected {
			t.Errorf("got %d, but expected %d", len(p.Contexts), expected)
		}
	})

	t.Run("verify that context is not existent in contexts", func(t *testing.T) {
		var expected Context = "stable"

		context, _ := p.existingContext(expected)
		if context {
			t.Errorf("context existing in contexts of project")
		}
	})
}

func projectWithSelectedContext(context Context) *Project {
	p := projectWithAddedContext("dev")
	p.SelectedContext = "dev"
	return p
}

func TestIsContextSelected(t *testing.T) {
	devSelected := projectWithSelectedContext("dev")
	devSelected.AddContext("int")

	t.Run("context is selected", func(t *testing.T) {
		var expected Context = "dev"

		if !devSelected.isContextSelected(expected) {
			t.Errorf("context '%v' is not the selected context", expected)
		}
	})

	t.Run("context is not selected", func(t *testing.T) {
		var expected Context = "int"

		if devSelected.isContextSelected(expected) {
			t.Errorf("context '%v' is the selected context", expected)
		}
	})
}

func TestUnselectContext(t *testing.T) {
	project := projectWithSelectedContext("dev")
	project.unselectContext()
	var expected Context = "unselected"

	if project.SelectedContext != expected {
		t.Errorf("got %v, but expected %v", project.SelectedContext, expected)
	}
}

func TestUnselectContextForAlreadyUnselectedContext(t *testing.T) {
	project := projectWithAddedContext("dev")
	project.unselectContext()
	var expected Context = "unselected"

	if project.SelectedContext != expected {
		t.Errorf("got %v, but expected %v", project.SelectedContext, expected)
	}
}

func TestSelectContext(t *testing.T) {
	project := projectWithMultipleContexts()
	project.SelectContext("stable")

	var expected Context = "stable"

	if project.SelectedContext != expected {
		t.Errorf("got %v, but expected %v", project.SelectedContext, expected)
	}
}

func TestSelectedContextNotPossibleIfContextNotExisting(t *testing.T) {
	project := projectWithMultipleContexts()
	expected := "given context: 'prod' not existing for project: 'business'"

	_, err := project.SelectContext("prod")

	if err.Error() != expected {
		t.Errorf("got \"%s\", but expected: \"%s\" ", expected, err.Error())
	}
}
