// Package translation provides functions to translate text into different languages.
package translation

import (
	"errors"
	"fmt"
	"sync"

	gtranslate "github.com/bas24/googletranslatefree"
)

// ErrInvalidLanguage is returned when an invalid language is provided.
var ErrInvalidLanguage = errors.New("invalid target language")

// TranslateText translates the given text into the target language.
func TranslateText(text string, from, to Language) (string, error) {
	if !IsValidLanguage(from) || !IsValidLanguage(to) {
		return "", fmt.Errorf("%w: %s", ErrInvalidLanguage, to)
	}

	translated, err := gtranslate.Translate(text, string(from), string(to))
	if err != nil {
		return "", fmt.Errorf("error translating text: %w", err)
	}

	return translated, nil
}

// BatchTranslate translates a batch of texts into the target language.
func BatchTranslate(texts []string, from, to Language) (map[string]string, error) {
	translations := make(map[string]string)
	var mu sync.Mutex
	var wg sync.WaitGroup
	var errs []error

	// If the language is not valid, return an error.
	if !IsValidLanguage(from) || !IsValidLanguage(to) {
		return translations, fmt.Errorf("%w: %s", ErrInvalidLanguage, to)
	}

	// Perform translations for valid languages.
	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			translated, err := TranslateText(text, from, to)
			mu.Lock()
			if err != nil {
				errs = append(errs, fmt.Errorf("error translating text '%s': %w", text, err))
				translations[text] = text // Fallback to original
			} else {
				translations[text] = translated
			}
			mu.Unlock()
		}(text)
	}

	wg.Wait()

	// If there were any errors during translation, return them as a batch.
	if len(errs) > 0 {
		return translations, fmt.Errorf("errors occurred during batch translation: %v", errs)
	}

	return translations, nil
}
