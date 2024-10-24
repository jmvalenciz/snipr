package model

type Snippet struct {
	Id     uint
	Name   string
	Note   string
	Body   string
	Tags   []string
	Format string
}

type CreateSnippet struct {
	Name   string
	Note   string
	Body   string
	Tags   []string
	Format string
}
