package models

type FilterItem struct {
	Label   string `json:"label"`
	Tooltip string `json:"tooltip"`
	Text    string `json:"text"`
}

type Filters struct {
	Columns [][]FilterItem `json:"columns"`
}
