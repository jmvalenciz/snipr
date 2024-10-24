package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type IView interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (IView, tea.Cmd, int)
	View() string
	OnRedraw() IView
}
