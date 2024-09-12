package tui

import (
	"baseconv/internal/converter"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type model struct {
	width int
	height int

	value converter.Value
	help help.Model
	keys keyMap
	
	inputs []textinput.Model
	selected int
}

const (
	INPUT_DEC = 0
	INPUT_HEX = 1
	INPUT_BIN = 2
	INPUT_OCT = 3
	INPUT_CHAR = 4
)

func NewModel() model {
	model := model {
		help: help.New(),
		keys: keys,
		inputs: make([]textinput.Model, 5),
	}

	for i, _ := range model.inputs {
		model.inputs[i] = textinput.New()
		model.inputs[i].Blur()

		switch i {
		case INPUT_DEC:
			model.inputs[i].Validate = converter.IsDecimal
			model.inputs[i].Prompt = "Dec: "
		case INPUT_HEX:
			model.inputs[i].Validate = converter.IsHexadecimal
			model.inputs[i].Prompt = "Hex: "
		case INPUT_BIN:
			model.inputs[i].Validate = converter.IsBinary
			model.inputs[i].Prompt = "Bin: "
		case INPUT_OCT:
			model.inputs[i].Validate = converter.IsOctal
			model.inputs[i].Prompt = "Oct: "
		case INPUT_CHAR:
			model.inputs[i].Validate = converter.IsCharacter
			model.inputs[i].Prompt = "Char: "
		}
	}
	model.inputs[INPUT_DEC].Focus()

	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		// cmd tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keys.Quit):
			if m.inputs[INPUT_CHAR].Focused() {
				break
			}

			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			m.selected = eumod(m.selected - 2, len(m.inputs))
		case key.Matches(msg, m.keys.Down):
			m.selected = eumod(m.selected + 2, len(m.inputs))
		case key.Matches(msg, m.keys.Left):
			m.selected = eumod(m.selected - 1, len(m.inputs))
		case key.Matches(msg, m.keys.Right):
			m.selected = eumod(m.selected + 1, len(m.inputs))
		}
	}

	cmds = append(cmds, m.updateInputs(msg))

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	style := lipgloss.NewStyle().Padding(1).Width(25)

	inputs := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Left, style.Render(m.inputs[INPUT_DEC].View()), style.Render(m.inputs[INPUT_HEX].View())),
		lipgloss.JoinHorizontal(lipgloss.Left, style.Render(m.inputs[INPUT_BIN].View()), style.Render(m.inputs[INPUT_OCT].View())),
		lipgloss.JoinHorizontal(lipgloss.Left, style.Render(m.inputs[INPUT_CHAR].View())),
	)
	b.WriteString(inputs + "\n\n")
	b.WriteString(leftPadding(1)(m.help.View(m.keys)))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Top,
		leftPadding(2)(b.String()),
	)
}
