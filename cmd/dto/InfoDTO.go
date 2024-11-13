package dto

import (
	"github.com/BurntSushi/toml"
)

type InfoData struct {
	Urls UrlsData `toml:"Urls"`
}

type UrlsData struct {
	Linkedin_url string `toml:"linkedin_url"`
	Github_url   string `toml:"github_url"`
}

func GetInfoData() *InfoData {
	var infoData InfoData
	toml.DecodeFile("./internal/data/info.toml", &infoData)

	return &infoData
}
