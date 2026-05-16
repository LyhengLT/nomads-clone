package models

type Meetup struct {
	Avatar    int    `json:"avatar"`
	Label     string `json:"label"`
	Rsvp      int    `json:"rsvp"`
	Attendees []int  `json:"attendees"`
}

type PopupFeature struct {
	Text   string  `json:"text"`
	Img    *string `json:"img"`
	ImgAlt *string `json:"imgAlt"`
}

type GenderOption struct {
	Label    string `json:"label"`
	Selected bool   `json:"selected"`
}

type PopupRight struct {
	Thumbnail        string         `json:"thumbnail"`
	MembersJoined    string         `json:"membersJoined"`
	EmailPlaceholder string         `json:"emailPlaceholder"`
	GenderOptions    []GenderOption `json:"genderOptions"`
	JoinBtn          string         `json:"joinBtn"`
	JoinHref         string         `json:"joinHref"`
}

type Meetups struct {
	Meetups       []Meetup       `json:"meetups"`
	Traveling     []int          `json:"traveling"`
	NewMembers    []int          `json:"newMembers"`
	PopupFeatures []PopupFeature `json:"popupFeatures"`
	PopupRight    PopupRight     `json:"popupRight"`
}
