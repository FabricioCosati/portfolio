package dto

type Language struct {
	Name    string
	Acronym string
	Image   string
}

type LanguagesCollection struct {
	Languages []Language
}

func GetSupportedLanguages() *LanguagesCollection {
	return &LanguagesCollection{
		Languages: []Language{
			{Name: "English", Acronym: "en", Image: "/img/us-flag.png"},
			{Name: "Português", Acronym: "pt", Image: "/img/pt-flag.png"},
		},
	}
}
