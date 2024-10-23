package services

import "errors"

var validLanguages = map[string]bool{
	"en": true,
	"pt": true,
}

func ValidateLanguage(lang string) error {
	if _, accept := validLanguages[lang]; !accept {
		return errors.New("language not valid")
	}

	return nil
}
