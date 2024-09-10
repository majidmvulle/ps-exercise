package models

type ContentWithLocalization struct {
	Content
	TextLocalized any `json:"-"`
}
