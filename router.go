// mavrw 20260210 2047CST
// freezing v1.0.0
package steep

type Router struct {
	stack  []*Screen
	cursor int
}

func NewRouter(root *Screen) *Router {
	r := &Router{
		stack: []*Screen{root},
	}

	return r
}

func (r *Router) Current() *Screen {
	return r.stack[r.cursor]
}

func (r *Router) Push(s *Screen) {
	// discard forward history
	r.stack = append(r.stack[:r.cursor+1], s)
	r.cursor++
}

func (r *Router) Replace(s *Screen) {
	r.stack[r.cursor] = s
}

func (r *Router) Close() {
	if r.cursor == 0 {
		return
	}
	r.stack = r.stack[:r.cursor]
	r.cursor--
}

func (r *Router) Back() {
	if !r.CanBack() {
		return
	}
	r.cursor--
}

func (r *Router) Forward() {
	if !r.CanForward() {
		return
	}
	r.cursor++
}

func (r *Router) CanBack() bool {
	return r.cursor > 0
}

func (r *Router) CanForward() bool {
	return r.cursor < len(r.stack)-1
}
