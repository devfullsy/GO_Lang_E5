package dictionary

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", fmt.Errorf("Le mot '%s' n'a pas été trouvé dans le dictionnaire", word)
	}
	return definition, nil
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() {

	var words []string
	for word := range d {
		words = append(words, word)
	}

	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s: %s\n", word, d[word])
	}
}
