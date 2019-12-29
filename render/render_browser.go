// +build js

package render

// RenderThread
// JavaScript only uses a single thread, therefore it's not required to forward calls to different threads.
// WebWorkers are currently not supported by this package.
type RenderThread struct {
}

// New returns a new render thread.
func New() *RenderThread {
	return &RenderThread{}
}

// Terminates the render thread (no-op).
func (r *RenderThread) Terminate() {
}

// Enqueue executes the given call.
// Every call is executed blocking.
func (r *RenderThread) Enqueue(blocking bool, fn func()) {
	_ = blocking
	fn()
}

// Sync synchronizes the render thread (no-op).
func (r *RenderThread) Sync() {
}
