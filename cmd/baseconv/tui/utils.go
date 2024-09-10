package tui

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func eumod(a, b int) int {
    if b == 0 {
        fmt.Fprintf(os.Stderr, "Error: Division by zero is undefined.\n");
        return 0;
    }
    
    var r = a % b
    if (r < 0) {
		if b > 0 {
			r += b
		} else {
			r += -b
		}
    }
    return r;
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	inputCmds := make([]tea.Cmd, len(m.inputs)+1)
	for i := 0; i < len(m.inputs); i++ {
		if i == m.selected {
			inputCmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focused
			m.inputs[i].TextStyle = focused
		} else {
			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = unfocused
			m.inputs[i].TextStyle = unfocused
		}

		if m.inputs[i].Err != nil {
			m.inputs[i].PromptStyle = error
			m.inputs[i].TextStyle = error
		}
	}
	cmds = append(cmds, tea.Batch(inputCmds...))


	for i := range m.inputs {
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	switch m.selected {
	case INPUT_DEC:
		val, err := strconv.ParseInt(m.inputs[INPUT_DEC].Value(), 10, 64)
		if err != nil {
			m.value.Reset()
		}

		m.value.UpdateDec(val)
	case INPUT_HEX:
		err := m.value.UpdateHex(m.inputs[INPUT_HEX].Value())
		if err != nil {
			m.value.Reset()
		}
	case INPUT_BIN:
		err := m.value.UpdateBin(m.inputs[INPUT_BIN].Value())
		if err != nil {
			m.value.Reset()
		}
	case INPUT_OCT:
		err := m.value.UpdateOct(m.inputs[INPUT_OCT].Value())
		if err != nil {
			m.value.Reset()
		}
	}
	
	m.inputs[INPUT_DEC].SetValue(fmt.Sprintf("%d", m.value.Dec))
	m.inputs[INPUT_HEX].SetValue(m.value.Hex)
	m.inputs[INPUT_BIN].SetValue(m.value.Bin)
	m.inputs[INPUT_OCT].SetValue(m.value.Oct)

	return tea.Batch(cmds...)
}
