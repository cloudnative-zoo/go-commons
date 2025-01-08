package translation

// Language represents a language code.
type Language string

// Define valid languages using iota for better maintainability.
const (
	English Language = "en"
	Urdu    Language = "ur"
	Swedish Language = "sv"
	Arabic  Language = "ar"
)

// validLanguages is a map of valid languages for quick lookup.
var validLanguages = map[Language]struct{}{
	English: {},
	Urdu:    {},
	Swedish: {},
	Arabic:  {},
}

// IsValidLanguage checks if the provided language is one of the allowed languages.
func IsValidLanguage(lang Language) bool {
	_, exists := validLanguages[lang]
	return exists
}
