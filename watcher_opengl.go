// +build !js

package gl

import (
	"github.com/go-gl/gl/v2.1/gl"
	"log"
)

// ContextWatcher is this library's context watcher, satisfying glfw.ContextWatcher interface.
// It must be notified when context is made current or detached.
var ContextWatcher = new(contextWatcher)

type contextWatcher struct {
	initGL bool
}

func (cw *contextWatcher) OnMakeCurrent(context interface{}) {
	if !cw.initGL {
		// Initialise gl bindings using the current context.
		err := gl.Init()
		if err != nil {
			log.Fatalln("gl.Init:", err)
		}
		cw.initGL = true
	}
}
func (contextWatcher) OnDetach() {}
