package view

import (
	"bytes"
	"fmt"
	"snipr/src/controller"
	"snipr/src/model"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SearchView struct {
	index      int
	snippets   []model.Snippet
	controller *controller.SnippetController
}

func NewSearchView(snippetController *controller.SnippetController) SearchView {
	var index int = -1
	snippets, err := snippetController.GetSnippets()
	if err != nil {
		panic(err)
	}
	if len(snippets) > 0 {
		index = 0
	}
	return SearchView{
		index:      index,
		snippets:   snippets,
		controller: snippetController,
	}
}

func (v SearchView) Init() tea.Cmd {
	return nil
}

func (v SearchView) Update(msg tea.Msg) (IView, tea.Cmd, int) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+d", "esc":
			return v, tea.Quit, -1
		case "up", "k":
			if v.index > 0 {
				v.index--
			}
		case "d":
			snippet := v.snippets[v.index]
			v.snippets = append(v.snippets[:v.index], v.snippets[v.index+1:]...)
			v.index--
			v.controller.DeleteById(snippet.Id)
		case "down", "j":
			if v.index < len(v.snippets)-1 {
				v.index++
			}
		case "c":
			snippet := v.snippets[v.index]
			clipboard.WriteAll(snippet.Body)
			return v, tea.Quit, -1
		case "a":
			return v, nil, CreateMode
		}
	}
	return v, nil, -1
}

func (v SearchView) View() string {
	layout := lipgloss.JoinHorizontal(lipgloss.Left, v.getSearchList(), v.getContent())
	return layout
}

func (v SearchView) OnRedraw() IView {
	var err error
	v.snippets, err = v.controller.GetSnippets()
	if err != nil {
		panic(err)
	}
	if len(v.snippets) > 0 {
		v.index = len(v.snippets) - 1
	} else {
		v.index = -1
	}
	return v
}

func (v SearchView) getContent() string {
	var snippet model.Snippet
	if v.index >= 0 {
		snippet = v.snippets[v.index]
	}
	style := lipgloss.NewStyle().MarginLeft(2).MarginTop(1).UnsetMaxWidth().UnsetMaxHeight().UnsetWidth().UnsetHeight()
	s := style.Render(lipgloss.JoinVertical(lipgloss.Top, "NAME: "+snippet.Name+" | FORMAT: "+snippet.Format+" | TAGS: "+v.getTags(snippet.Tags), "\n"+v.getNote(snippet.Note), "\n"+v.getBody(snippet.Body, snippet.Format)))
	return s
}

func (v SearchView) getTags(tags []string) string {
	s := ""
	styled_tags := make([]string, len(tags))
	for i, tag := range tags {
		styled_tags[i] = tagStyles[i%len(tagStyles)].Render(tag)
	}
	s += strings.Join(styled_tags, ",")
	return s
}

func (v SearchView) getBody(body string, format string) string {
	buf := new(bytes.Buffer)
	quick.Highlight(buf, body, format, "terminal", "monokai")
	s := buf.String()
	return s
}

func (v SearchView) getNote(note string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	return style.Render(note)
}

func (v SearchView) getSearchList() string {
	s := "Search List\n"
	style := lipgloss.NewStyle().BorderRight(true).BorderStyle(lipgloss.NormalBorder()).Width(24)
	for i, snippet := range v.snippets {
		if i == v.index {
			s += colorPrimary.Render(fmt.Sprintf("\n> %s", snippet.Name))
		} else {
			s += fmt.Sprintf("\n  %s", snippet.Name)
		}
		// s += lipgloss.NewStyle().Foreground(lipgloss.Color(5)).Render()
	}
	return style.Render(s)
}
