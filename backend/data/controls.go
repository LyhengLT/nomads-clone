package data

import "backend/models"

var GetControls = models.Controls{
	JoinBtn:         "Join Nomads.com →",
	CompareBtn:      "Compare",
	GridView:        []string{"Grid View", "Map View", "Chart View"},
	GridViewDefault: "Grid View",
	NomadScore: []models.NomadScoreGroup{
		{
			Group:   "General",
			Options: []string{"Nomad Score", "Solomad score", "Family score", "Quality of life score", "Your Score", "Cost for nomad", "Internet Speed", "Fun", "Safety"},
		},
		{
			Group:   "Social",
			Options: []string{"Members rating", "Nomads there now", "Nomads been"},
		},
		{
			Group:   "Cost",
			Options: []string{"Cost for nomad", "Cost for local", "Cost for expat"},
		},
	},
	NomadScoreDefault: "Nomad Score",
}
