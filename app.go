// mavrw 20260210 2117CST
// freezing v1.0.0
package steep

import tea "github.com/charmbracelet/bubbletea"

type App struct {
	router *Router
}

func NewApp(root *Screen) *App {
	return &App{
		router: NewRouter(root),
	}
}

// Init delegates to the active screen's Init.
func (a *App) Init() tea.Cmd {
	return a.router.Current().Init()
}

// Update delegates to the active screen's Update, and
// also handles navigation messages automatically.
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// handle nvigation messages first
	switch msg := msg.(type) {
	case PushMsg:
		a.router.Push(msg.Screen)
	case ReplaceMsg:
		a.router.Replace(msg.Screen)
	case CloseMsg:
		a.router.Close()
	case BackMsg:
		a.router.Back()
	case ForwardMsg:
		a.router.Forward()
	}

	// delegate to the current screen
	m, cmd := a.router.Current().Update(msg)
	if m != nil {
		return m, cmd
	}
	return a, cmd
}

// View delegates to the active screen's View.
func (a *App) View() string {
	return a.router.Current().View()
}

// Router returns the underlying Router in the event
// the developer wants to inspect or manipulate it directly.
func (a *App) Router() *Router {
	return a.router
}
