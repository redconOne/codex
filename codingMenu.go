package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func codingMenu() *huh.Form {
	problemTypes := huh.NewOptions(
		"Arrays & Hashing",
		"Two Pointers",
		"Stack",
		"Binary Search",
		"Sliding Window",
		"Linked List",
		"Trees",
		"Tries",
		"Backtracking",
		"Heap & Priority Queue",
		"Graphs",
		"1-D DP",
		"Intervals",
		"Greedy",
		"Advanced Graphs",
		"2-D DP",
		"Bit Manipulation",
		"Math & Geometry",
	)
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("problemType").
				Options(problemTypes...).
				Title("What would you like to do today?").
				Description("Choose an topic to practice please"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)
}

func codingMenuUpdate(m *Model, msg tea.Msg) ([]tea.Cmd, error) {
	var cmds []tea.Cmd
	form, cmd := m.codingMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.codingMenu = f
		cmds = append(cmds, cmd)
	} else {
		return cmds, fmt.Errorf("error updating codingMenu")
	}

	if m.codingMenu.State == huh.StateCompleted {
		m.menuType = "completed"
		cmds = append(cmds, tea.Quit)
	}

	return cmds, nil
}

func codingMenuView(m *Model, s *Styles) string {
	v := strings.TrimSuffix(m.codingMenu.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	errors := m.codingMenu.Errors()
	header := m.appBoundaryView("Codex")
	if len(errors) > 0 {
		header = m.appErrorBoundaryView(m.errorView())
	}
	body := lipgloss.JoinHorizontal(lipgloss.Top, form)

	defaultHelpCommands := m.codingMenu.Help().ShortHelpView(m.codingMenu.KeyBinds())
	quitCommands := m.styles.Help.Render(" • esc to quit")
	fullCommands := defaultHelpCommands + quitCommands

	footer := m.appBoundaryView(fullCommands)
	if len(errors) > 0 {
		footer = m.appErrorBoundaryView("")
	}

	return s.Base.Render(header + "\n" + body + "\n\n" + footer)
}
