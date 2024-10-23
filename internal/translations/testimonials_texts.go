package translations

import "github.com/nicksnyder/go-i18n/v2/i18n"

func LoadTestimonialsTexts() map[string]i18n.Message {
	return map[string]i18n.Message{
		"Testimonial01Date": {
			ID:    "Testimonial01Date",
			Other: "On October 19, 2024, Rafael was working with Fabrício on the same team",
		},
		"Testimonial01Text01": {
			ID:    "Testimonial01Text01",
			Other: "“It is a privilege to speak about the experience I had working with Fabrício. He is always in search of the best solution to problems, focusing on making the code as performant and readable as possible. With extensive knowledge in backend development, especially in API integrations with payment systems such as credit cards and PIX payments, he taught me a lot.",
		},
		"Testimonial01Text02": {
			ID:    "Testimonial01Text02",
			Other: "In addition to mastering Python, Golang, Laravel, PHP, React, and AWS, Fabrício has a solid understanding of SOLID programming principles. Always willing to help the team, he constantly seeks the best strategies for each task.”",
		},
		"Testimonial02Date": {
			ID:    "Testimonial02Date",
			Other: "On October 18, 2024, Ryan was working with Fabrício on the same team",
		},
		"Testimonial02Text01": {
			ID:    "Testimonial02Text01",
			Other: "“I had the pleasure of working with Fabrício on the technology team at Celcoin, and I can say it was an exceptional experience. His focus and responsibility toward deliveries and innovations are remarkable, with an admirable dedication to teamwork. He has strong skills in Clean Code, TDD, design patterns, Golang, PHP/Laravel, and React, as well as proficiency in cloud tools like AWS. His expertise in refactoring, high performance, scalability, and automated testing truly stands out and is widely admired by all of us. Fabrício's dedication and his constant pursuit of excellence make him an exceptional professional on our team.”",
		},
		"Testimonial03Date": {
			ID:    "Testimonial03Date",
			Other: "On October 17, 2024, Tales was working with Fabrício on the same team",
		},
		"Testimonial03Text01": {
			ID:    "Testimonial03Text01",
			Other: "“I had the opportunity to work with Fabrício at Celcoin on the Cash Management front. During this time, I was able to closely observe his work. Fabrício is an extremely competent and helpful professional. He pays great attention to detail and raises important points during refinements. He excels at solving complex problems, which enhances collaboration and helps the team move forward more efficiently.",
		},
		"Testimonial03Text02": {
			ID:    "Testimonial03Text02",
			Other: "Working with him was a rewarding experience, and I am certain that he will make a difference on any team he is part of.”",
		},
	}
}
