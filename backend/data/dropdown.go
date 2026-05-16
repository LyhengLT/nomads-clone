package data

import "backend/models"

var GetDropdown = []models.DropdownSection{
	{
		Type: "top",
		Links: []models.DropdownLink{
			{Href: "#", Text: "👩‍💻 Join Remote OK"},
			{Href: "#", Text: "👋 Log in"},
		},
	},
	{
		Section: "GENERAL",
		Links: []models.DropdownLink{
			{Href: "#", Text: "🌐 Frontpage"},
			{Href: "#", Text: "🌗 Dark mode"},
			{Href: "#", Text: "❤️ Your favorites"},
			{Href: "#", Text: "🚑 Nomad insurance"},
		},
	},
	{
		Section: "COMMUNITY",
		Links: []models.DropdownLink{
			{Href: "#", Text: "❤️ Dating app", Badge: "new"},
			{Href: "#", Text: "🤝 Friend finder"},
			{Href: "#", Text: "💬 Chat"},
			{Href: "#", Text: "🌎 Members map"},
			{Href: "#", Text: "💈 Host meetup"},
			{Href: "#", Text: "🍹 Attend meetup"},
		},
	},
	{
		Section: "TOOLS",
		Links: []models.DropdownLink{
			{Href: "#", Text: "🏜️ Explore"},
			{Href: "#", Text: "🎒 NomadGPT"},
			{Href: "#", Text: "📸 Vote on photos"},
			{Href: "#", Text: "💸 FIRE calculator"},
			{Href: "#", Text: "💥 Fastest growing"},
			{Href: "#", Text: "🧪 State of remote work", Badge: "new"},
			{Href: "#", Text: "🕸️ Network graph"},
			{Href: "#", Text: "🏆 Best place now"},
			{Href: "#", Text: "🔮 Random good place"},
			{Href: "#", Text: "🗓️ Residence calendar"},
			{Href: "#", Text: "⛅️ Climate finder"},
			{Href: "#", Text: "📊 Nomad stats", Badge: "new"},
			{Href: "#", Text: "📜 History of nomads", Badge: "new"},
			{Href: "#", Text: "🔌 Fastest internet"},
			{Href: "#", Text: "🔮 Random place"},
		},
	},
	{
		Section: "FEEDS",
		Links: []models.DropdownLink{
			{Href: "#", Text: "🛠 Remote Jobs API"},
			{Href: "#", Text: "🪚 RSS feed"},
			{Href: "#", Text: "🪓 JSON feed"},
			{Href: "#", Text: "🔶 Hacker News mode"},
			{Href: "#", Text: "💙 Safe for work mode"},
		},
	},
	{
		Section: "HELP",
		Links: []models.DropdownLink{
			{Href: "#", Text: "💡 Ideas + bugs"},
			{Href: "#", Text: "🚀 Changelog"},
			{Href: "#", Text: "🛍️ Merch"},
			{Href: "#", Text: "🛟 FAQ & Help"},
		},
	},
	{
		Section: "OTHER PROJECTS",
		Links: []models.DropdownLink{
			{Href: "#", Text: "📊 Remote work stats", Badge: "new"},
			{Href: "#", Text: "👷 Top remote companies", Badge: "new"},
			{Href: "#", Text: "💰 Highest paying remote jobs", Badge: "new"},
			{Href: "#", Text: "🌍 Become a digital nomad"},
			{Href: "#", Text: "🔮 Web3 Jobs"},
			{Href: "#", Text: "📸 Photo AI", Badge: "new"},
			{Href: "#", Text: "🏡 Interior AI"},
		},
	},
}
