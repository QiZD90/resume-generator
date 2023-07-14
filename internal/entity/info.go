package entity

type ResumeInfo struct {
	Name       string
	Profession string
	Age        string

	Image []byte
	Crop  Crop

	Education    []string
	Achievements []string
	Languages    []string
	Skills       []string
	Contacts     []Contact

	Columns int
}

type Contact struct {
	Name     string
	Link     string
	LinkText string `json:"link_text"`
	Text     string
}

type Crop struct {
	X, Y, Size int
}
