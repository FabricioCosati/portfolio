package translations

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	Bundle i18n.Bundle
}

func InitTranslator() (*Translator, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if _, err := bundle.LoadMessageFile("./internal/translations/active.en.toml"); err != nil {
		return nil, err
	}

	if _, err := bundle.LoadMessageFile("./internal/translations/active.pt.toml"); err != nil {
		return nil, err
	}

	return &Translator{
		Bundle: *bundle,
	}, nil
}

func (t *Translator) InitLocalize(language string) *i18n.Localizer {
	return i18n.NewLocalizer(&t.Bundle, language)
}

func (t *Translator) GetTranslatedMessage(localizer i18n.Localizer, id string) (string, error) {
	return localizer.LocalizeMessage(&i18n.Message{ID: id})
}
