package models

type Todo struct {
	ID          int64  `json:"id"` //when be in json rename to "id"
	Title       string `json:"title"`
	Description string `jstron:"description"`
	Done        bool   `josn:"done"`
}
