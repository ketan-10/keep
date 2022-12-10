// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package entities

type NewNote struct {
	Text        string `json:"text"`
	Description string `json:"description"`
}

type Note struct {
	ID          int        `json:"id"`
	Text        string     `json:"text"`
	Description string     `json:"description"`
	SubNote     []*SubNote `json:"subNote"`
}

type SubNote struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
}
