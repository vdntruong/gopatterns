package main

import "log"

func main() {
	var searchTitle = "Go"
	pm := CreateProcessManager()
	result := pm.Find(ByTitle(searchTitle))

	if len(result) == 0 {
		log.Println("No results found")
		return
	}
	log.Println("Result:", result)
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
