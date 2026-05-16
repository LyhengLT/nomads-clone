package models

type SidebarItem struct {
	Text      string `json:"text"`
	Tooltip   string `json:"tooltip,omitempty"`
	FullWidth bool   `json:"fullWidth,omitempty"`
}

type SidebarRow struct {
	Row   string        `json:"row"`
	Items []SidebarItem `json:"items"`
}

type SidebarWhereRow struct {
	Items []string `json:"items"`
}

type SidebarWhenRow struct {
	Items []string `json:"items"`
	Stick bool     `json:"stick"`
}

type SidebarScoreRow struct {
	Items []string `json:"items"`
}

type SidebarHowRow struct {
	Items []string `json:"items"`
	Stick bool     `json:"stick"`
}

type SidebarPassport struct {
	Placeholder string `json:"placeholder"`
}

type Sidebar struct {
	What            []SidebarRow      `json:"what"`
	WhereGrid       []string          `json:"whereGrid"`
	WhereRows       []SidebarWhereRow `json:"whereRows"`
	WhenRows        []SidebarWhenRow  `json:"whenRows"`
	Months          []string          `json:"months"`
	Score           []SidebarScoreRow `json:"score"`
	Passports       []SidebarPassport `json:"passports"`
	PassportOptions []string          `json:"passportOptions"`
	How             []SidebarHowRow   `json:"how"`
}
