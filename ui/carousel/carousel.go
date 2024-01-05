package carousel

// A from-scratch carousel bubble

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	listColor         lipgloss.Color = lipgloss.Color("#343238")
	selectedListColor lipgloss.Color = lipgloss.Color("#a947b3")

	listStyle lipgloss.Style = lipgloss.NewStyle().
			Width(10).Align(lipgloss.Center)
)

type Model struct {
	Length    int
	listStart int
	List      []string
	Index     int
	focused   bool
}

func New() Model {
	return Model{
		Length:    3,
		Index:     0,
		List:      []string{},
		listStart: 0,
		focused:   false,
	}
}

func (m *Model) Focus() {
	m.focused = true
}

func (m *Model) Blur() {
	m.focused = false
}

func (m *Model) Value() string {
	return m.List[m.Index]
}

func (m *Model) SetValue(text string) {
	m.List[m.Index] = text
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focused {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.Index = max(0, m.Index-1)
			if m.Index < m.listStart {
				m.listStart--
			}
		case "right":
			m.Index = min(len(m.List)-1, m.Index+1)
			if m.Index > m.listStart+m.Length-1 {
				m.listStart++
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var view strings.Builder
	var style lipgloss.Style = listStyle.Copy()

	view.WriteString(style.Width(3).Render("<"))
	style.Background(listColor).Width(10)

	width := min(m.Length, len(m.List))

	right := min(len(m.List), width+m.listStart)

	for i, planet := range m.List[m.listStart:right] {

		if i+m.listStart == m.Index {
			style.Background(selectedListColor)
			view.WriteString(style.Render(planet))
			style.Background(listColor)
		} else {
			view.WriteString(style.Render(planet))
		}

	}
	style.UnsetBackground().Width(3)
	view.WriteString(style.Render(">"))

	return view.String()
}
