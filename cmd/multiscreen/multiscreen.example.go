package main

import (
	"fmt"
	"os"

	"github.com/mavrw/steep"

	tea "github.com/charmbracelet/bubbletea"
)

// screens
var (
	home = steep.NewScreen(
		"home",
		nil, // default Init
		func(msg tea.Msg) (tea.Model, tea.Cmd) {
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "enter": // go to next screen
					return nil, steep.PushCmd(step1)
				case "q", "ctrl": // quit
					return nil, tea.Quit
				}
			}
			return nil, nil
		},
		steep.ViewFromString("Home Screen\nPress 'enter' for Step 1, 'q' to quit"),
	)

	step1 = steep.NewScreen(
		"step1",
		nil,
		func(msg tea.Msg) (tea.Model, tea.Cmd) {
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "backspace": // go back to home
					return nil, steep.BackCmd()
				case "enter": // push step 2
					return nil, steep.PushCmd(step2)
				case "q", "ctrl": // quit
					return nil, tea.Quit
				}
			}
			return nil, nil
		},
		steep.ViewFromString("Step 1 Screen\nPress 'backspace' to go back, 'enter' for Step 2, 'q' to quit"),
	)

	step2 = steep.NewScreen(
		"step2",
		nil,
		func(msg tea.Msg) (tea.Model, tea.Cmd) {
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "backspace": // back to step1
					return nil, steep.BackCmd()
				case "q", "ctrl": // quit
					return nil, tea.Quit
				}
			}
			return nil, nil
		},
		steep.ViewFromString("Step 2 Screen\nPress 'backspace' to go back, 'q' to quit"),
	)
)

func main() {
	// create App with root screen
	app := steep.NewApp(home)

	// run Bubble Tea program
	p := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
