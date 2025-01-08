package main

import (
	"fmt"
	"log"

	"github.com/cloudnative-zoo/go-commons/translation"
)

func main() {
	texts := []string{"Hello", "World"}
	toLang := translation.Urdu

	translations, err := translation.BatchTranslate(texts, toLang)
	if err != nil {
		log.Printf("Translation errors: %v", err)
	}

	for original, translated := range translations {
		fmt.Printf("Original: %s, Translated: %s\n", original, translated)
	}
}
