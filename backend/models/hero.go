package models

type HeroBenefit struct {
	Emoji string `json:"emoji"`
	Text  string `json:"text"`
	Href  string `json:"href"`
}

type HeroNewsLogo struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type HeroQuote struct {
	Quote      string `json:"quote"`
	Img        string `json:"img"`
	Alt        string `json:"alt"`
	ExtraClass string `json:"extraClass"`
}

type HeroLaurel struct {
	Text  string `json:"text"`
	Stars int    `json:"stars"`
	Since string `json:"since"`
	Img   string `json:"img"`
}

type HeroCtaBox struct {
	Thumbnail        string `json:"thumbnail"`
	EmailPlaceholder string `json:"emailPlaceholder"`
	JoinBtn          string `json:"joinBtn"`
	AlreadyAccount   string `json:"alreadyAccount"`
}

type Hero struct {
	Profiles     []string       `json:"profiles"`
	Benefits     []HeroBenefit  `json:"benefits"`
	NewsLogos    []HeroNewsLogo `json:"newsLogos"`
	Quotes       []HeroQuote    `json:"quotes"`
	Laurel       HeroLaurel     `json:"laurel"`
	HeroTitle    string         `json:"heroTitle"`
	HeroSubtitle string         `json:"heroSubtitle"`
	CtaBox       HeroCtaBox     `json:"ctaBox"`
}
