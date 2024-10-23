package dto

type Language struct {
	Name    string
	Acronym string
	Image   string
}

type LanguagesCollection struct {
	Languages []Language
}
