package translations

import "github.com/nicksnyder/go-i18n/v2/i18n"

func LoadHomeTexts() map[string]i18n.Message {
	return map[string]i18n.Message{
		"Welcome": {
			ID:    "Welcome",
			Other: "Hello there! I'm",
		},
		"FirstPartSubtitle": {
			ID:    "FirstPartSubtitle",
			Other: "A self-taught",
		},
		"SecondPartSubtitle": {
			ID:    "SecondPartSubtitle",
			Other: "Software Developer",
		},
		"ThirdPartSubtitle": {
			ID:    "ThirdPartSubtitle",
			Other: "with a great interest in technology",
		},
		"Resume": {
			ID:    "Resume",
			Other: "See my Resume",
		},
		"About": {
			ID:    "About",
			Other: "About Me",
		},
		"AboutFirstParagraph": {
			ID:    "AboutFirstParagraph",
			Other: "I am a developer with a strong interest in technology and a focus on creating simple solutions for complex problems.",
		},
		"AboutSecondParagraph": {
			ID:    "AboutSecondParagraph",
			Other: "Currently, I am part of the integrations team at Celcoin, in the BU cel_cash, where I work on APIs focusing on payment and receipt on PIX, credit card collections, and financial integrations.",
		},
		"AboutThirdParagraph": {
			ID:    "AboutThirdParagraph",
			Other: "I always seek to keep myself updated and enjoy sharing knowledge and learning from others.",
		},
		"MyJourney": {
			ID:    "MyJourney",
			Other: "My Journey",
		},
		"CelcoinPosition": {
			ID:    "CelcoinPosition",
			Other: "Full Stack Developer Junior",
		},
		"EscalarPosition": {
			ID:    "EscalarPosition",
			Other: "Backend Developer, Intern",
		},
		"CollegeDegree": {
			ID:    "CollegeDegree",
			Other: "Bachelor's Degree, Information Systems",
		},
		"CollegeName": {
			ID:    "CollegeName",
			Other: "Una University Center",
		},
		"PresentTime": {
			ID:    "PresentTime",
			Other: "Present",
		},
		"Location": {
			ID:    "Location",
			Other: "Belo Horizonte, Brazil",
		},
		"Skills": {
			ID:    "Skills",
			Other: "Skills",
		},
		"FrontendDevelopment": {
			ID:    "FrontendDevelopment",
			Other: "Frontend Development",
		},
		"BackendDevelopment": {
			ID:    "BackendDevelopment",
			Other: "Backend Development",
		},
		"CloudInfra": {
			ID:    "CloudInfra",
			Other: "Cloud Infra-Architecture",
		},
		"Testimonials": {
			ID:    "Testimonials",
			Other: "Testimonials",
		},
	}
}
