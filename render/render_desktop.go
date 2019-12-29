// +build !js

package render

import "runtime"

const callBufSize = 20

// RenderThread is locked to a OS thread and receives async and/or sync commands to execute.
type RenderThread struct {
	calls    chan call
	finished chan struct{}
}

type call struct {
	fn   func()
	done chan<- struct{}
}

// New creates and starts a new render thread.
func New() *RenderThread {
	r := &RenderThread{
		calls:    make(chan call, callBufSize),
		finished: make(chan struct{}),
	}
	go r.run()
	return r
}

func (r *RenderThread) run() {
	runtime.LockOSThread()
	defer close(r.finished)
	for c := range r.calls {
		c.fn()
		if c.done != nil {
			close(c.done)
		}
	}
}

// Terminate the render thread.
// The thread will complete the remaining, queued calls and stop afterwards.
// This function blocks until the render thread completed.
// No new calls can be enqueued.
func (r *RenderThread) Terminate() {
	close(r.calls)
	<-r.finished
}

// Enqueue a call to be executed by the render thread.
// If blocking is true, the function call will wait until the render thread executed it.
// If false, the call is async.
func (r *RenderThread) Enqueue(blocking bool, fn func()) {
	var done chan struct{}
	if blocking {
		done = make(chan struct{})
	}
	r.calls <- call{
		fn:   fn,
		done: done,
	}
	if blocking {
		<-done
	}
}

// Sync synchronizes the render thread.
// Waits until the render thread completed all commands that were enqueued until now.
// It is possible to issue new commands while waiting for a synchronization point.
func (r *RenderThread) Sync() {
	r.Enqueue(true, func() {})
}
