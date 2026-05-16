package models

type NomadScoreGroup struct {
	Group   string   `json:"group"`
	Options []string `json:"options"`
}

type Controls struct {
	JoinBtn           string            `json:"joinBtn"`
	CompareBtn        string            `json:"compareBtn"`
	GridView          []string          `json:"gridView"`
	GridViewDefault   string            `json:"gridViewDefault"`
	NomadScore        []NomadScoreGroup `json:"nomadScore"`
	NomadScoreDefault string            `json:"nomadScoreDefault"`
}
