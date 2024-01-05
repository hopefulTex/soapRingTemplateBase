package ui

import (
	"fmt"
	"soapRingTest/ui/carousel"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var Planets []string = []string{
	"Mercury",
	"Venus",
	"Earth",
	"Mars",
	"Jupiter",
	"Saturn",
	"Uranus",
	"Neptune",
	"[REDACTED]",
}

type Model struct {
	planetID int

	keys KeyMap
	help help.Model

	list carousel.Model

	textinput textinput.Model

	Width  int
	Height int

	editing  bool
	quitting bool
}

func NewProgram() *tea.Program {
	// func NewProgram(cfg Config) *tea.Program {
	// normalize options from flags / cobra

	model := newModel()
	// model := newModel(options)
	return tea.NewProgram(model)
}

// depends on args handler
func newModel() Model {
	// func newModel(options Options) {
	t := textinput.New()
	c := carousel.New()
	c.List = Planets
	c.Focus()
	//t.Placeholder = Planets[0]
	return Model{
		planetID:  0,
		keys:      DefaultKeyMap,
		help:      help.New(),
		Width:     40,
		Height:    24,
		list:      c,
		textinput: t,
		editing:   false,
		quitting:  false,
	}
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, textinput.Blink)

	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.help.Width = m.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, DefaultKeyMap.Edit):
			if m.editing {
				m.list.SetValue(m.textinput.Value())
				m.textinput.Reset()
				m.textinput.Blur()
				m.list.Focus()
			} else {
				m.textinput.Focus()
				m.textinput.SetValue(m.list.Value())
				m.list.Blur()
			}
			m.editing = !m.editing
		default:
			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)
			m.list, cmd = m.list.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var view strings.Builder

	view.WriteString(m.list.View())
	view.WriteRune('\n')
	view.WriteString(helloWorld(m.list.Value()))

	str := mainStyle.Render(view.String())

	view.Reset()
	view.WriteString(str)
	if m.editing {
		view.WriteRune('\n')
		view.WriteString(m.textinput.View())
	}

	if m.quitting {
		return fmt.Sprintf("%s\nGoodbye, %s!\n", str, m.list.Value())
	}

	view.WriteRune('\n')
	view.WriteString(m.help.View(m.keys))

	return view.String()
}

func helloWorld(planet string) string {
	return fmt.Sprintf("Hello, %s!\n", planet)
}
