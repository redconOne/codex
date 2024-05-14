package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 80

var (
	red                   = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo                = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green                 = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	termWidth, termHeight = getTerminalSize()
	marginLeft            = termWidth / 3
	marginTop             = termHeight / 3
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
		Padding(1, 4, 0, 1).
		Margin(marginTop, 0, 0, marginLeft)
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

func getTerminalSize() (width, height int) {
	fd := uintptr(syscall.Stdout)

	var dimensions [4]uint16
	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0); errno != 0 {
		return 0, 0
	}

	width = int(dimensions[1])
	height = int(dimensions[0])

	return width, height
}

// TODO: add support for additional submenus (Testing, Stats, Options)
func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.menuType = "mainMenu"

	m.mainMenu = mainMenu()
	m.codingMenu = codingMenu()

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
	var err error

	// TODO: add cases for additional submenus (Testing, Stats, Options)
	switch m.menuType {
	case "mainMenu":
		cmds, err = mainMenuUpdate(&m, msg)
	case "codingMenu":
		cmds, err = codingMenuUpdate(&m, msg)
	}

	if err != nil {
		log.Fatalf("error updating menu: %s", err)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	// TODO: add cases for additional submenus (Testing, Stats, Options)
	switch m.menuType {
	case "codingMenu":
		return codingMenuView(&m, s)
	case "mainMenu":
		return mainMenuView(&m, s)
	case "completed":
		var b strings.Builder
		fmt.Fprintf(&b, "Generating template for problem #%s\n\n", "1234")
		fmt.Fprintf(&b, "Please standby...")
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	case "exit":
		var b strings.Builder
		fmt.Fprintf(&b, "Thanks for using Codex!")
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:
		var b strings.Builder
		fmt.Fprintf(&b, "This area is still under construction\n\n")
		fmt.Fprintf(&b, "Content here soon...")
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
