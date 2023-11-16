// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	dictionary "projet1/directory/directories"
)

func main() {
	// Charger les données depuis le fichier JSON
	data, err := ioutil.ReadFile("donnees.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON:", err)
		return
	}

	// Définir la structure pour stocker les données du fichier JSON
	var jsonData struct {
		Data []map[string]string `json:"data"`
	}

	// Décoder les données JSON dans la structure
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	// Initialiser le dictionnaire avec les données du fichier JSON
	myDictionary := make(dictionary.Dictionary)
	for _, item := range jsonData.Data {
		for word, definition := range item {
			myDictionary.Add(word, definition)
		}
	}

	// Utiliser le dictionnaire
	definition, err := myDictionary.Get("go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Définition de 'go': %s\n", definition)
	}

	myDictionary.Remove("python")

	fmt.Println("Liste des mots et définitions:")
	myDictionary.List()
}
