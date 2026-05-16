package models

type CityWeather struct {
	Icon    string `json:"icon"`
	Feels   int    `json:"feels"`
	Temp    int    `json:"temp"`
	Humid   string `json:"humid"`
	Aqi     int    `json:"aqi"`
	AqiIcon string `json:"aqiIcon"`
}

type CityBars struct {
	Overall  string `json:"overall"`
	Cost     string `json:"cost"`
	Internet string `json:"internet"`
	Liked    string `json:"liked"`
	Safety   string `json:"safety"`
}

type City struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Country string      `json:"country"`
	Img     string      `json:"img"`
	Rank    int         `json:"rank"`
	Wifi    int         `json:"wifi"`
	Popular bool        `json:"popular"`
	Weather CityWeather `json:"weather"`
	Cost    string      `json:"cost"`
	Bars    CityBars    `json:"bars"`
	Type    string      `json:"type"`
}
