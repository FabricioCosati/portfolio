package services

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fabricio-cosati/portfolio/internal/translations"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/allegro/bigcache"
)

type LocalizerTranslator struct {
	Localizer i18n.Localizer
	Language  string
}

var cache *bigcache.BigCache

func InitCache() {
	var err error
	cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))

	if err != nil {
		log.Fatalf("failed to initialize cache: %v", err)
	}
}

func InitLocalize(localizer i18n.Localizer, language string) LocalizerTranslator {
	return LocalizerTranslator{
		Localizer: localizer,
		Language:  language,
	}
}

func GetCachedText(key string, loader func() map[string]string) (map[string]string, error) {
	entry, err := cache.Get(key)

	if err == nil {
		var data map[string]string
		json.Unmarshal(entry, &data)
		return data, nil
	}

	data := loader()
	jsonData, _ := json.Marshal(data)
	cache.Set(key, jsonData)

	return data, nil
}

func (l *LocalizerTranslator) LoadHomeTexts() (map[string]string, error) {
	return GetCachedText("home_texts_"+l.Language, func() map[string]string {
		return l.GetTranslatedTexts(translations.LoadHomeTexts())
	})
}

func (l *LocalizerTranslator) LoadTestimonialsTexts() (map[string]string, error) {
	return GetCachedText("testimonials_text_"+l.Language, func() map[string]string {
		return l.GetTranslatedTexts(translations.LoadTestimonialsTexts())
	})
}

func (l *LocalizerTranslator) GetTranslatedTexts(texts map[string]i18n.Message) map[string]string {
	translated := make(map[string]string)

	for _, text := range texts {
		translated_text, _ := l.Localizer.LocalizeMessage(&i18n.Message{ID: text.ID})

		translated[text.ID] = translated_text
	}
	return translated
}
