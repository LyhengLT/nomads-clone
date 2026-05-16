package data

import "backend/models"

var GetHero = models.Hero{
	HeroTitle:    "🌍 Go nomad",
	HeroSubtitle: "Join a global community of remote workers living and traveling around the world",
	Profiles: []string{
		"images/profiles/profile-1.png",
		"images/profiles/profile-2.png",
		"images/profiles/profile-3.png",
		"images/profiles/profile-4.png",
		"images/profiles/profile-5.png",
		"images/profiles/profile-6.png",
		"images/profiles/profile-7.png",
		"images/profiles/profile-8.png",
		"images/profiles/profile-9.png",
		"images/profiles/profile-10.png",
		"images/profiles/profile-11.png",
	},
	Benefits: []models.HeroBenefit{
		{Emoji: "🍹", Text: "Attend 277 meetups/year in 100+ cities", Href: "#"},
		{Emoji: "❤️", Text: "Meet new people for dating and friends", Href: "#"},
		{Emoji: "🧪", Text: "Research destinations and find your best place to live and work", Href: "#"},
		{Emoji: "🌎", Text: "Keep track of your travels and record where you've been", Href: "#"},
		{Emoji: "💬", Text: "Join Nomads.com chat and find your community on the road", Href: "#"},
	},
	NewsLogos: []models.HeroNewsLogo{
		{Src: "images/news/nyt.png", Alt: "New York Times"},
		{Src: "images/news/ft.png", Alt: "Financial Times"},
		{Src: "images/news/bbc.png", Alt: "BBC"},
		{Src: "images/news/cnn.png", Alt: "CNN"},
		{Src: "images/news/usa-today.png", Alt: "USA Today"},
		{Src: "images/news/cnbc.png", Alt: "CNBC"},
		{Src: "images/news/guardian.png", Alt: "Guardian"},
		{Src: "images/news/politico.png", Alt: "Politico"},
	},
	Quotes: []models.HeroQuote{
		{
			Quote:      "\"Nomads.com ranks destinations that are accommodating to digital nomads, based on factors like cost of living, internet speed and weather\"",
			Img:        "images/news/nyt.png",
			Alt:        "New York Times",
			ExtraClass: "",
		},
		{
			Quote:      "\"The rankings of Nomads.com's cities are constantly in flux (all the data is refreshed in real-time based on user input)\"",
			Img:        "images/news/bbc.png",
			Alt:        "BBC",
			ExtraClass: "center-card",
		},
		{
			Quote:      "\"Nomads.com is a Kayak-like aggregator for potential work destinations, ranking internet, price, and safety\"",
			Img:        "images/news/techcrunch.png",
			Alt:        "TechCrunch",
			ExtraClass: "",
		},
	},
	Laurel: models.HeroLaurel{
		Text:  "#1 Nomad Community",
		Stars: 5,
		Since: "Since 2014",
		Img:   "images/laurel.svg",
	},
	CtaBox: models.HeroCtaBox{
		Thumbnail:        "images/thumbnail.jpg",
		EmailPlaceholder: "Type your email...",
		JoinBtn:          "Join Nomads.com →",
		AlreadyAccount:   "If you already have an account, we'll log you in",
	},
}
