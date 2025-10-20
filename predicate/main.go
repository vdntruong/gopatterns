package main

import "log"

func main() {
	pm := CreateProcessManager()

	title := "Go"
	result := pm.Find(ByTitle(title))

	if len(result) == 0 {
		log.Println("No results found")
		return
	}
	log.Println("Result:", result)
}
