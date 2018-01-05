package models

type Tag struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Total    int    `json:"total"`
	Unreaded int    `json:"unreaded"`
}

type TaggedItem struct {
	Id      int    `json:"id"`
	TagId   int    `json:"tagid"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
	Link    string `json:"link"`
	Date    int64  `json:"date"`
	Source  int    `json:"source"` // 1 stack
}
