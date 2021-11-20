package db

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	{
		Title:   "Hello",
		Desc:    "Article Description",
		Content: "Article Content",
	},
	{
		Title:   "Hello2",
		Desc:    "Article Description",
		Content: "Article Content",
	},
}
