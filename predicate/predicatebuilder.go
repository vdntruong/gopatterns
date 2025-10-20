package main

type Process struct {
	ID    int
	Title string
}

type ProcessPredicate func(*Process) bool

func ByTitle(title string) ProcessPredicate {
	return func(p *Process) bool {
		return p.Title == title
	}
}

func ByID(id int) ProcessPredicate {
	return func(p *Process) bool {
		return p.ID == id
	}
}
