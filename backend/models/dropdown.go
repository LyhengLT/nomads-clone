package models

type DropdownLink struct {
	Href  string `json:"href"`
	Text  string `json:"text"`
	Badge string `json:"badge"`
}

type DropdownSection struct {
	Type    string         `json:"type"`
	Section string         `json:"section"`
	Links   []DropdownLink `json:"links"`
}
