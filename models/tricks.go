package models

type Trick struct {
	Id            string          `json:"id"`
	Name          string          `json:"name"`
	Aliases       []string        `json:"aliases"`
	Categories    []TrickCategory `json:"categories"`
	Prerequisites []string        `json:"prereqs"`
	NextTricks    []string        `json:"nextTricks"`
	Description   string          `json:"description"`
}
