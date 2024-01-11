package dictionary

import (
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type Dictionary struct {
	words         map[string]string
	addChannel    chan addOperation
	removeChannel chan removeOperation
}

type addOperation struct {
	word, definition string
}

type removeOperation struct {
	word string
}

func NewDictionary() *Dictionary {
	d := &Dictionary{
		words:         make(map[string]string),
		addChannel:    make(chan addOperation),
		removeChannel: make(chan removeOperation),
	}

	go d.processOperations()
	return d
}

func (d *Dictionary) processOperations() {
	for {
		select {
		case addOp := <-d.addChannel:
			d.words[addOp.word] = addOp.definition
		case removeOp := <-d.removeChannel:
			delete(d.words, removeOp.word)
		}
	}
}

func (d *Dictionary) Add(word, definition string) error {

	if err := validateData(word, definition); err != nil {
		return &ValidationError{Message: err.Error()}
	}

	d.addChannel <- addOperation{word, definition}
	return nil
}

func validateData(word, definition string) error {
	minLength := 3
	maxLength := 50

	if len(word) < minLength || len(word) > maxLength {
		return fmt.Errorf("La longueur du mot doit être entre %d et %d caractères", minLength, maxLength)
	}

	if len(definition) < minLength || len(definition) > maxLength {
		return fmt.Errorf("La longueur de la définition doit être entre %d et %d caractères", minLength, maxLength)
	}

	return nil
}

func (d *Dictionary) Remove(word string) {
	d.removeChannel <- removeOperation{word}
}

func (d *Dictionary) Get(word string) (string, error) {
	definition, ok := d.words[word]
	if !ok {
		return "", &NotFoundError{Message: fmt.Sprintf("Le mot '%s' n'a pas été trouvé dans le dictionnaire", word)}
	}
	return definition, nil
}

func (d *Dictionary) List() string {
	var words []string
	for word := range d.words {
		words = append(words, word)
	}

	sort.Strings(words)

	var result strings.Builder
	for _, word := range words {
		result.WriteString(fmt.Sprintf("%s: %s\n", word, d.words[word]))
	}

	return result.String()
}
