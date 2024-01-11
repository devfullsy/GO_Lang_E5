package dictionary

import (
	"fmt"
	"sort"
)

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

func (d *Dictionary) Add(word, definition string) {
	d.addChannel <- addOperation{word, definition}
}

func (d *Dictionary) Remove(word string) {
	d.removeChannel <- removeOperation{word}
}

func (d *Dictionary) Get(word string) (string, error) {
	definition, ok := d.words[word]
	if !ok {
		return "", fmt.Errorf("Le mot '%s' n'a pas été trouvé dans le dictionnaire", word)
	}
	return definition, nil
}

func (d *Dictionary) List() {
	var words []string
	for word := range d.words {
		words = append(words, word)
	}

	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s: %s\n", word, d.words[word])
	}
}
