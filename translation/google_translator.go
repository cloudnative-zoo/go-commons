// Package translation provides functions to translate text into different languages.
package translation

import (
	"errors"
	"fmt"
	"sync"

	gtranslate "github.com/gilang-as/google-translate"
)

// ErrInvalidLanguage is returned when an invalid language is provided.
var ErrInvalidLanguage = errors.New("invalid target language")

// TranslateText translates the given text into the target language.
func TranslateText(text string, toLang Language) (string, error) {
	if !IsValidLanguage(toLang) {
		return "", fmt.Errorf("%w: %s", ErrInvalidLanguage, toLang)
	}

	value := gtranslate.Translate{
		Text: text,
		To:   string(toLang),
	}

	translated, err := gtranslate.Translator(value)
	if err != nil {
		return "", fmt.Errorf("error translating text: %w", err)
	}

	return translated.Text, nil
}

// BatchTranslate translates a batch of texts into the target language.
func BatchTranslate(texts []string, toLang Language) (map[string]string, error) {
	translations := make(map[string]string)
	var mu sync.Mutex
	var wg sync.WaitGroup
	var errs []error

	// If the language is not valid, return an error.
	if !IsValidLanguage(toLang) {
		return nil, fmt.Errorf("%w: %s", ErrInvalidLanguage, toLang)
	}

	// Perform translations for valid languages.
	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			translated, err := TranslateText(text, toLang)
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
