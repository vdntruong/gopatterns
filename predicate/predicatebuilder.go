package main

import "fmt"

type Process struct {
	ID    int
	Title string
}

func (p *Process) String() string {
	return fmt.Sprintf("{ID: %d, Title: %s}", p.ID, p.Title)
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

type ProcessManager struct {
	processes []*Process
}

func CreateProcessManager() *ProcessManager {
	return &ProcessManager{
		// default processes
		processes: []*Process{
			{ID: 1, Title: "Go"},
			{ID: 2, Title: "Python"},
			{ID: 3, Title: "C++"},
		},
	}
}

func (pm *ProcessManager) Find(predicate ProcessPredicate) []*Process {
	var result []*Process
	for _, p := range pm.processes {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}
