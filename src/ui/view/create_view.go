package view

import (
	"snipr/src/controller"
	"snipr/src/model"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CreateView struct {
	newSnippet   model.CreateSnippet
	controller   *controller.SnippetController
	nameInput    textinput.Model
	bodyInput    textarea.Model
	formatInput  textinput.Model
	noteInput    textarea.Model
	tagsInput    textinput.Model
	currentInput int
}

func NewCreateView(snippetController *controller.SnippetController) CreateView {
	nameInput := textinput.New()
	bodyInput := textarea.New()
	bodyInput.ShowLineNumbers = true
	bodyInput.SetHeight(8)
	bodyInput.MaxWidth = 32
	bodyInput.CharLimit = 0
	formatInput := textinput.New()
	noteInput := textarea.New()
	noteInput.ShowLineNumbers = true
	noteInput.SetHeight(3)
	tagsInput := textinput.New()
	tagsInput.Placeholder = "js,dotenv,yaml"
	nameInput.Focus()

	return CreateView{
		controller:   snippetController,
		newSnippet:   model.CreateSnippet{},
		nameInput:    nameInput,
		bodyInput:    bodyInput,
		formatInput:  formatInput,
		noteInput:    noteInput,
		tagsInput:    tagsInput,
		currentInput: 0,
	}
}

func (v CreateView) Init() tea.Cmd {
	return v.blinkInput()
}

func (v CreateView) Update(msg tea.Msg) (IView, tea.Cmd, int) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return v, nil, SearchMode
		case "ctrl+c", "ctrl+d":
			return v, tea.Quit, -1
		case "ctrl+p":
			if v.currentInput > 0 {
				v.blurInput()
				v.currentInput--
				v.focusInput()
			}
		case "ctrl+n":
			if v.currentInput < 4 {
				v.blurInput()
				v.currentInput++
				v.focusInput()
			} else {
				newSnippet := model.CreateSnippet{
					Name:   v.nameInput.Value(),
					Body:   v.bodyInput.Value(),
					Format: v.formatInput.Value(),
					Note:   v.noteInput.Value(),
					Tags:   strings.Split(v.tagsInput.Value(), ","),
				}
				_, err := v.controller.Create(newSnippet)
				if err != nil {
					panic(err)
				}
				return v, nil, SearchMode
			}
		}
	case error:
		panic(msg)
	}
	view, cmd := v.updateInput(msg)
	return view, cmd, -1
}

func (v CreateView) View() string {
	inputs := v.nameInput.View()
	inputs += v.bodyInput.View()
	return lipgloss.JoinVertical(
		lipgloss.Top,
		"\nName:",
		v.nameInput.View(),
		"\nBody:",
		v.bodyInput.View(),
		"\nFormat:",
		v.formatInput.View(),
		"\nNote:",
		v.noteInput.View(),
		"\nTags:",
		v.tagsInput.View(),
		v.getHelp(),
	)
}

func (v CreateView) OnRedraw() IView {
	v.nameInput.SetValue("")
	v.bodyInput.SetValue("")
	v.formatInput.SetValue("")
	v.noteInput.SetValue("")
	v.tagsInput.SetValue("")
	v.blurInput()
	v.currentInput = 0
	v.focusInput()
	return v
}

func (v CreateView) getHelp() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	return style.Render("\n<TAB> next | <SHIFT>+<TAB> previous | <CTRL>+n create")
}

func (v CreateView) blinkInput() tea.Cmd {
	switch v.currentInput {
	case 0:
		return v.nameInput.Cursor.BlinkCmd()
	case 1:
		return v.bodyInput.Cursor.BlinkCmd()
	case 2:
		return v.formatInput.Cursor.BlinkCmd()
	case 3:
		return v.noteInput.Cursor.BlinkCmd()
	case 4:
		return v.tagsInput.Cursor.BlinkCmd()
	default:
		return nil
	}
}

func (v *CreateView) blurInput() {
	switch v.currentInput {
	case 0:
		v.nameInput.Blur()
	case 1:
		// v.bodyInput.Blur()
		// v.bodyInput.KeyMap.InsertNewline.SetEnabled(false)
	case 2:
		v.formatInput.Blur()
	case 3:
		// v.noteInput.Blur()
		v.noteInput.KeyMap.InsertNewline.SetEnabled(false)
	case 4:
		v.tagsInput.Blur()
	default:
	}
}

func (v CreateView) updateInput(msg tea.Msg) (CreateView, tea.Cmd) {
	var cmd tea.Cmd
	switch v.currentInput {
	case 0:
		v.nameInput, cmd = v.nameInput.Update(msg)
		return v, cmd
	case 1:
		v.bodyInput, cmd = v.bodyInput.Update(msg)
		return v, cmd
	case 2:
		v.formatInput, cmd = v.formatInput.Update(msg)
		return v, cmd
	case 3:
		v.noteInput, cmd = v.noteInput.Update(msg)
		return v, cmd
	case 4:
		v.tagsInput, cmd = v.tagsInput.Update(msg)
		return v, cmd
	default:
		panic("")
	}
}

func (v *CreateView) focusInput() tea.Cmd {
	switch v.currentInput {
	case 0:
		return v.nameInput.Focus()
	case 1:
		return v.bodyInput.Focus()
	case 2:
		return v.formatInput.Focus()
	case 3:
		return v.noteInput.Focus()
	case 4:
		return v.tagsInput.Focus()
	default:
		return nil
	}
}
