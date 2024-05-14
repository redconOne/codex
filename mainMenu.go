package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func mainMenu() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("mainChoice").
				Options(huh.NewOptions("Start coding", "Test solution", "View stats", "Configure settings", "About Codex", "Exit")...).
				Title("What would you like to do today?").
				Description("Choose an operation please"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)
}

func mainMenuUpdate(m *Model, msg tea.Msg) ([]tea.Cmd, error) {
	var cmds []tea.Cmd
	form, cmd := m.mainMenu.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.mainMenu = f
		cmds = append(cmds, cmd)
	} else {
		return cmds, fmt.Errorf("error updating mainMenu")
	}

	if m.mainMenu.State == huh.StateCompleted {
		mainChoice := m.mainMenu.GetString("mainChoice")
		switch mainChoice {
		case "Start coding":
			m.menuType = "codingMenu"
		case "Exit":
			m.menuType = "exit"
			cmds = append(cmds, tea.Quit)
		default:
			m.menuType = "other"
			cmds = append(cmds, tea.Quit)
		}
	}

	return cmds, nil
}

func mainMenuView(m *Model, s *Styles) string {
	v := strings.TrimSuffix(m.mainMenu.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	errors := m.mainMenu.Errors()
	header := m.appBoundaryView("Codex")
	if len(errors) > 0 {
		header = m.appErrorBoundaryView(m.errorView())
	}
	body := lipgloss.JoinHorizontal(lipgloss.Top, form)

	defaultHelpCommands := m.mainMenu.Help().ShortHelpView(m.mainMenu.KeyBinds())
	quitCommands := m.styles.Help.Render(" • esc to quit")
	fullCommands := defaultHelpCommands + quitCommands

	footer := m.appBoundaryView(fullCommands)
	if len(errors) > 0 {
		footer = m.appErrorBoundaryView("")
	}

	return s.Base.Render(header + "\n" + body + "\n\n" + footer)
}
