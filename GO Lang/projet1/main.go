package main

import (
	"fmt"
	dictionary "projet1/directory/directories"
)

func main() {

	myDictionary := make(dictionary.Dictionary)

	myDictionary.Add("go", "language de programmation")
	myDictionary.Add("mathématique", "science exacte")
	myDictionary.Add("mma", "sport de combat")

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
