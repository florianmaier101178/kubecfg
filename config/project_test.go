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

func projectWithMultipleContexts() *Project {
	p := projectWithAddedContext("dev")
	p.AddContext("int")
	p.AddContext("stable")
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

func TestAddContextForMultipleContextsExistent(t *testing.T) {
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

func TestRemoveContextNotPossibleIfNoMatchingContextExisting(t *testing.T) {
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

func TestRemoveContextWithMultipleContextsExistent(t *testing.T) {
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
