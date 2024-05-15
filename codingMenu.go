package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func codingMenu() *huh.Form {
	category := "categories"
	arr := []*huh.Group{}

	topicGroup := huh.NewGroup(
		huh.NewSelect[string]().
			Title("What topic would you like to practice?").
			Value(&category).
			Options(
				huh.NewOptions(topicList...)...,
			),
	).WithHideFunc(func() bool { return category != "categories" })

	arr = append(arr, topicGroup)

	for key, list := range problemLists {
		newGroup := huh.NewGroup(
			huh.NewSelect[string]().
				Title("What problem would you like to practice?").
				Value(&category).
				Options(
					huh.NewOptions(list...)...,
				),
		).WithHideFunc(func() bool { return category != key })

		arr = append(arr, newGroup)
	}

	return huh.NewForm(
		arr...,
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

func codingMenuView(m *Model) string {
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

	return lipgloss.Place(
		termWidth,
		termHeight,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Top,
			header,
			body,
			footer,
		),
	)
}
