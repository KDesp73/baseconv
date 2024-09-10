package tui

import "github.com/charmbracelet/lipgloss"

var (
	focused = lipgloss.NewStyle().Foreground(lipgloss.Color("57"))
	unfocused = lipgloss.NewStyle().Foreground(lipgloss.Color("#f0f0f0"))
	error = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))
	leftPadding = func(size int) func(strs ...string) string {
		return lipgloss.NewStyle().PaddingLeft(size).Render
	}
)
