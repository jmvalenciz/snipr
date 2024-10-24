package view

import "github.com/charmbracelet/lipgloss"

var (
	colorPrimary = lipgloss.NewStyle().Foreground(lipgloss.Color("13"))
	tagStyles    = []lipgloss.Style{
		lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("2")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("3")),
		lipgloss.NewStyle().Foreground(lipgloss.Color("4")),
	}
)
