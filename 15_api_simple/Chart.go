package main

//ChartResponse ... the model for api response
type Maintainer struct {
	Name  string
	Email string
}
type Source struct {
	Self string
}
type Repo struct {
	Name string
	URL  string
}
type Attribute struct {
	Name        string
	Repo        Repo
	Description string
	Home        string
	keywords    []string
	Maintainers []Maintainer
	Sources     []Source
	Icon        string
}

type Link struct {
	Self string
}

type Records struct {
	ID         string
	Type       string
	Attributes Attribute
	Link       Link
}

type Meta struct {
	totalPages int
}

type ChartResponse struct {
	data []Records
	meta Meta `json:"meta"`
}
