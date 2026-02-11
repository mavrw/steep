package steep

// mavrw 20260210 2002CST
// freezing v1.0.0
import tea "github.com/charmbracelet/bubbletea"

var _ tea.Model = (*Screen)(nil)

type ScreenID string

type screenInitFunc func() tea.Cmd
type screenUpdateFunc func(msg tea.Msg) (tea.Model, tea.Cmd)
type screenViewFunc func() string

// Screen wraps three functions:
// Init, Update, and View; thus implementing tea.Model
type Screen struct {
	id       ScreenID
	initFn   screenInitFunc
	updateFn screenUpdateFunc
	viewFn   screenViewFunc
}

// NewScreen returns a Screen that delegates Init, Update, and View
// to the functions provided in the call. Nil functions default to no-ops
func NewScreen(id ScreenID,
	ifn screenInitFunc, ufn screenUpdateFunc, vfn screenViewFunc) *Screen {
	if ifn == nil {
		ifn = NoOpInit
	}
	if ufn == nil {
		ufn = NoOpUpdate
	}
	if vfn == nil {
		vfn = NoOpView
	}

	return &Screen{
		id:       id,
		initFn:   ifn,
		updateFn: ufn,
		viewFn:   vfn,
	}
}

// Init implements tea.Model.
func (s *Screen) Init() tea.Cmd {
	return s.initFn()
}

// Update implements tea.Model.
func (s *Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s.updateFn(msg)
}

// View implements tea.Model.
func (s *Screen) View() string {
	return s.viewFn()
}

// ID returns the screen's identifier.
func (s *Screen) ID() ScreenID {
	return s.id
}

// NoOpInit returns nil.
func NoOpInit() tea.Cmd { return nil }

// NoOpUpdate returns (nil, nil).
// NoOpUpdate should only be used when the router ignores the returned model
func NoOpUpdate(msg tea.Msg) (tea.Model, tea.Cmd) { return nil, nil }

// NoOpView returns an empty string.
func NoOpView() string { return "" }

// InitFromCmd returns a screenInitFunc.
func InitFromCmd(cmd tea.Cmd) screenInitFunc {
	return func() tea.Cmd { return cmd }
}

// ViewFromString returns a screenViewFunc that renders s.
func ViewFromString(s string) screenViewFunc {
	return func() string { return s }
}
