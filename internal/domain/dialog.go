package domain

type DialogData struct {
	Dialogs []Dialog `json:"dialogs"`
}

type Dialog struct {
	Slug    string   `json:"slug"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Text     string `json:"text"`
	GoToSlug string `json:"go_to_slug"`
}
