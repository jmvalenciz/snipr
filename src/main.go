package main

import (
	"log"
	"snipr/src/cmd"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	logFile, err := tea.LogToFile("./debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %w", err)
	}
	cmd.Execute()
	defer logFile.Close()
}
