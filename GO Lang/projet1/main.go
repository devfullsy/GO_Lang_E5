package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	dictionary "projet1/directory/directories"
	"sync"
)

func main() {
	data, err := ioutil.ReadFile("donnees.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON:", err)
		return
	}

	var jsonData struct {
		Data []map[string]string `json:"data"`
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	myDictionary := dictionary.NewDictionary()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, item := range jsonData.Data {
			for word, definition := range item {
				myDictionary.Add(word, definition)
			}
		}
	}()

	go func() {
		defer wg.Done()
		myDictionary.Remove("mathématique")
	}()

	wg.Wait()

	definition, err := myDictionary.Get("go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Définition de 'go': %s\n", definition)
	}

	fmt.Println("Liste des mots et définitions:")
	myDictionary.List()
}
