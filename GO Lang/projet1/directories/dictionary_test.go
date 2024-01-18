package dictionary

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	d := NewDictionary()

	err := d.Add("test", "definition")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	definition, err := d.Get("test")
	if err != nil || definition != "definition" {
		t.Errorf("Expected definition 'definition', got '%s'", definition)
	}
}

func TestAddValidationFailure(t *testing.T) {
	d := NewDictionary()

	err := d.Add("ab", "definition")
	if err == nil || err.Error() != "La longueur du mot doit être entre 3 et 50 caractères" {
		t.Errorf("Expected validation error, got %v", err)
	}

	err = d.Add("test", "a"+strings.Repeat("b", 51))
	if err == nil || err.Error() != "La longueur de la définition doit être entre 3 et 50 caractères" {
		t.Errorf("Expected validation error, got %v", err)
	}
}

func TestRemove(t *testing.T) {
	d := NewDictionary()

	d.Add("test", "definition")
	d.Remove("test")

	_, err := d.Get("test")
	if err == nil {
		t.Error("Expected error, got none")
	}
}

func TestGetNotFound(t *testing.T) {
	d := NewDictionary()

	_, err := d.Get("nonexistent")
	if err == nil || err.Error() != "Le mot 'nonexistent' n'a pas été trouvé dans le dictionnaire" {
		t.Errorf("Expected not found error, got %v", err)
	}
}

func TestList(t *testing.T) {
	d := NewDictionary()

	d.Add("test1", "definition1")
	d.Add("test2", "definition2")
	d.Add("test3", "definition3")

	expected := "test1: definition1\ntest2: definition2\ntest3: definition3\n"
	result := d.List()

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}
