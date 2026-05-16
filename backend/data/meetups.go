package data

import "backend/models"

func strPtr(s string) *string { return &s }

var GetMeetups = models.Meetups{
	Meetups: []models.Meetup{
		{Avatar: 1, Label: "Wed 18th Mar: Istanbul", Rsvp: 1, Attendees: []int{}},
		{Avatar: 2, Label: "Wed 18th Mar: Da Nang", Rsvp: 7, Attendees: []int{1, 2, 3, 4, 5}},
		{Avatar: 3, Label: "Thu 12th Mar: Luque", Rsvp: 2, Attendees: []int{6, 7}},
	},
	Traveling:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5},
	NewMembers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5},
	PopupFeatures: []models.PopupFeature{
		{Text: "<b>Join a global community</b> of remote workers living and traveling around the world", Img: strPtr("images/join-modal-community.jpg"), ImgAlt: strPtr("join modal")},
		{Text: "<b>Get unlimited members-only access to cities in 195+ countries</b> and research the best places to live, work and travel remotely", Img: strPtr("images/join-modal-data.jpg"), ImgAlt: strPtr("join modal")},
		{Text: "<b>Get access to the paid Nomads.com Chat on Telegram</b> and find your community on the road (15,000+ messages sent this month)", Img: strPtr("images/join-modal-telegram.jpg"), ImgAlt: strPtr("nomad chat")},
		{Text: "<b>Meet new and interesting people</b> for dating or friends and get matched based on your interests", Img: strPtr("images/join-modal-dating.jpg"), ImgAlt: strPtr("nomad chat")},
		{Text: "<b>Attend Nomads.com meetups in cities all over the world</b> and make friends (4 meetups per week, 208 per year)", Img: strPtr("images/meeting.jpg"), ImgAlt: strPtr("meeting")},
		{Text: "<b>Keep track of your travels</b> to record where you've been, where you're going and meet people there", Img: nil, ImgAlt: nil},
		{Text: "<b>Unlock hundreds of filters</b> to pick destinations on exactly your preferences and find your perfect place", Img: nil, ImgAlt: nil},
		{Text: "<b>Learn how to get visas and residence permits</b> from other people who been through the process", Img: strPtr("images/passport.jpg"), ImgAlt: strPtr("passport")},
		{Text: "<b>Get alerts when you're approaching</b> the 183 days when you become a legal/tax resident in a country", Img: nil, ImgAlt: nil},
		{Text: "<b>Use the Climate Finder</b> to search cities based on future weather and air quality", Img: nil, ImgAlt: nil},
		{Text: "<b>Use the FIRE retirement tool</b> to find places to retire early", Img: nil, ImgAlt: nil},
		{Text: "Awarded <b>Best Remote Work Tool</b> and Best Travel <b>App</b>", Img: strPtr("images/join-modal-ph-award.jpg"), ImgAlt: strPtr("ph award")},
		{Text: "<b>Be the first to receive new remote jobs</b> on our Nomads.com Chat", Img: nil, ImgAlt: nil},
		{Text: "<b>Get access to Airline List Pro</b> to filter airlines with data on best safety, service and live lost luggage data", Img: nil, ImgAlt: nil},
	},
	PopupRight: models.PopupRight{
		Thumbnail:        "images/thumbnail.jpg",
		MembersJoined:    "🔥 653 members joined this month",
		EmailPlaceholder: "Type your email...",
		GenderOptions: []models.GenderOption{
			{Label: "Select your gender", Selected: true},
			{Label: "Man", Selected: false},
			{Label: "Woman", Selected: false},
			{Label: "Non-binary", Selected: false},
			{Label: "Prefer not to say", Selected: false},
		},
		JoinBtn:  "Join Nomads.com →",
		JoinHref: "#",
	},
}
