package gl

var ctx *context

type RenderThread interface {
	Enqueue(blocking bool, fn func())
}

type context struct {
	enqueue func(blocking bool, fn func())
}

// Init initializes the OpenGL package.
// Expects a render thread to execute OpenGL commands.
func Init(renderThread RenderThread) {
	if ctx != nil {
		panic("already initialized")
	}
	ctx = &context{
		enqueue: renderThread.Enqueue,
	}
}
