package data

import "backend/models"

var GetSidebar = models.Sidebar{
	What: []models.SidebarRow{
		{Row: "three", Items: []models.SidebarItem{
			{Text: "🍦 Cold Now", Tooltip: "Filter by cold weather"},
			{Text: "🌤️ Mild now", Tooltip: "Filter by mild weather"},
			{Text: "☀️ Warm now", Tooltip: "Filter by warm weather"},
		}},
		{Row: "three", Items: []models.SidebarItem{
			{Text: "💵 &lt;US$1k/mo", Tooltip: "Costs < $1,000/mo"},
			{Text: "💸 &lt;US$2k/mo", Tooltip: "Costs < $2,000/mo"},
			{Text: "💰 &lt;US$3k/mo", Tooltip: "Costs < $3,000/mo"},
		}},
		{Row: "two", Items: []models.SidebarItem{
			{Text: "👮‍ Safe", Tooltip: ""},
			{Text: "📡 Fast internet", Tooltip: ""},
		}},
		{Row: "two", Items: []models.SidebarItem{
			{Text: "💨 Clean air now", Tooltip: ""},
			{Text: "👍 Liked by members", Tooltip: ""},
		}},
		{Row: "two", Items: []models.SidebarItem{
			{Text: "🔥 Popular now", Tooltip: ""},
			{Text: "📈 Growing in nomads", Tooltip: ""},
		}},
		{Row: "two stick", Items: []models.SidebarItem{
			{Text: "🏅 Top ranked", Tooltip: ""},
			{Text: "💎 Hidden gem", Tooltip: ""},
		}},
		{Row: "two", Items: []models.SidebarItem{
			{Text: "✨ You haven't been", Tooltip: ""},
			{Text: "✨ For you", Tooltip: ""},
		}},
		{Row: "two", Items: []models.SidebarItem{
			{Text: "🇰🇭 Near Cambodia", Tooltip: "", FullWidth: true},
		}},
	},
	WhereGrid: []string{
		"🌎 North America",
		"💃 Latin America",
		"🇪🇺 Europe",
		"🌍 Africa",
		"🕌 Middle East",
		"⛩️ Asia",
		"🏄 Oceania",
		"🛰️ Space",
	},
	WhereRows: []models.SidebarWhereRow{
		{Items: []string{"🇺🇸 United States", "🏝 Caribbean"}},
		{Items: []string{"🇪🇺 European Union", "🇪🇺 Not in Schengen"}},
		{Items: []string{"🤝 Easy to make friends", "❤️ Great for dating"}},
	},
	WhenRows: []models.SidebarWhenRow{
		{Items: []string{"☃️ In the winter", "♻️ All year round"}, Stick: true},
	},
	Months: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
	Score: []models.SidebarScoreRow{
		{Items: []string{"🤩 Exceptional (4.75+)", "😍 Very good (4.5+)"}},
		{Items: []string{"👍 Good (4.25+)", "😐 Okay (3+)"}},
	},
	Passports: []models.SidebarPassport{
		{Placeholder: "Select your passport"},
		{Placeholder: "Select your partner's passport"},
		{Placeholder: "Select your friend's passport"},
	},
	PassportOptions: []string{
		"🇰🇭 Cambodia",
		"🇺🇸 United States",
		"🇬🇧 United Kingdom",
		"🇩🇪 Germany",
		"🇫🇷 France",
		"🇯🇵 Japan",
		"🇸🇬 Singapore",
		"🇦🇺 Australia",
	},
	How: []models.SidebarHowRow{
		{Items: []string{"👮‍ Not Humid", "📡 Not rainy now"}, Stick: false},
		{Items: []string{"💨 Clean air now", "👍 Liked by members"}, Stick: false},
		{Items: []string{"🔥 Popular now", "📈 Growing in nomads"}, Stick: false},
		{Items: []string{"🏅 Top ranked", "💎 Hidden gem"}, Stick: true},
		{Items: []string{"✨ You haven't been", "✨ For you"}, Stick: false},
	},
}
