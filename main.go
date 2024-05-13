package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 80

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

type state int

// TODO: add support for additional submenus (Testing, Stats, Options)
type Model struct {
	lg         *lipgloss.Renderer
	styles     *Styles
	mainMenu   *huh.Form
	codingMenu *huh.Form
	menuType   string
	state      state
	width      int
}

// TODO: add support for additional submenus (Testing, Stats, Options)
func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.menuType = "mainMenu"

	m.mainMenu = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("mainChoice").
				Options(huh.NewOptions("Start coding", "Test solution", "View stats", "Configure settings", "About Codex", "Exit")...).
				Title("What would you like to do today?").
				Description("Choose an operation please"),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)

	m.codingMenu = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("problemType").
				Options(huh.NewOptions("Array/Hashing", "Sliding Window", "Linked List", "Greedy", "Stack")...).
				Title("What would you like to do today?").
				Description("Choose an operation please"),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)

	return m
}

func (m Model) Init() tea.Cmd {
	return m.mainMenu.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// TODO: add cases for additional submenus (Testing, Stats, Options)
	switch m.menuType {
	case "mainMenu":
		form, cmd := m.mainMenu.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.mainMenu = f
			cmds = append(cmds, cmd)
		}

		if m.mainMenu.State == huh.StateCompleted {
			mainChoice := m.mainMenu.GetString("mainChoice")
			switch mainChoice {
			case "Start coding":
				m.menuType = "codingMenu"
			default:
				cmds = append(cmds, tea.Quit)
			}
		}
	case "codingMenu":
		form, cmd := m.codingMenu.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.codingMenu = f
			cmds = append(cmds, cmd)
		}

		if m.codingMenu.State == huh.StateCompleted {
			m.menuType = "completed"
			cmds = append(cmds, tea.Quit)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	// TODO: add cases for additional submenus (Testing, Stats, Options)
	switch m.menuType {
	case "codingMenu":
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

	case "mainMenu":
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

	default:
		problemType := m.codingMenu.GetString("problemType")
		problemType = s.Highlight.Render(problemType)
		var b strings.Builder
		fmt.Fprintf(&b, "You chose to work on: %s\n", problemType)
		fmt.Fprintf(&b, "Please wait while templates are generated...")
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	}
}

func (m Model) errorView() string {
	var s string
	for _, err := range m.mainMenu.Errors() {
		s += err.Error()
	}
	return s
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

func main() {
	_, err := tea.NewProgram(NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
