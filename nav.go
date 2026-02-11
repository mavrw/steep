// mavrw 20260210 2116CST
// freezing v1.0.0
package steep

import tea "github.com/charmbracelet/bubbletea"

// messages
type PushMsg struct {
	Screen *Screen
}

type ReplaceMsg struct {
	Screen *Screen
}

type CloseMsg struct{}
type BackMsg struct{}
type ForwardMsg struct{}

// helper commands
func PushCmd(s *Screen) tea.Cmd    { return func() tea.Msg { return PushMsg{Screen: s} } }
func ReplaceCmd(s *Screen) tea.Cmd { return func() tea.Msg { return ReplaceMsg{Screen: s} } }
func CloseCmd() tea.Cmd            { return func() tea.Msg { return CloseMsg{} } }
func BackCmd() tea.Cmd             { return func() tea.Msg { return BackMsg{} } }
func ForwardCmd() tea.Cmd          { return func() tea.Msg { return ForwardMsg{} } }
