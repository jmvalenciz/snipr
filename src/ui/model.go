package ui

import (
	"snipr/src/controller"
	view "snipr/src/ui/view"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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

type Model struct {
	views       map[int]view.IView
	currentView int
}

func NewModel(snippetController *controller.SnippetController) Model {
	views := make(map[int]view.IView)
	currentView := view.SearchMode
	views[view.SearchMode] = view.NewSearchView(snippetController)
	views[view.CreateMode] = view.NewCreateView(snippetController)
	return Model{
		currentView: currentView,
		views:       views,
	}
}

func (m Model) Init() tea.Cmd {
	return m.views[m.currentView].Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	view := m.views[m.currentView]
	updatedView, cmd, newMode := view.Update(msg)
	m.views[m.currentView] = updatedView
	if newMode != -1 {
		m.currentView = newMode
		m.views[newMode] = m.views[newMode].OnRedraw()
	}
	return m, cmd
}

func (m Model) View() string {
	return m.views[m.currentView].View()
}
