# gl

Package gl is a Go cross-platform binding for OpenGL, with an OpenGL ES 2-like API.

It supports:

- **macOS**, **Linux** and **Windows** via OpenGL 2.1 backend,

- **iOS** and **Android** via OpenGL ES 2.0 backend,

- **Modern Browsers** (desktop and mobile) via WebGL 1.0 backend.

It was forked of https://github.com/goxjs/gl with [#28](https://github.com/goxjs/gl/issues/28) resolved
to add concurrency support similar to http://golang.org/x/mobile/gl, from where the project originally stems. 

The main differences to http://golang.org/x/mobile/gl are the added support for web ([CL 8793](https://go-review.googlesource.com/8793)),
as well as the focus on being a lightweight library with good desktop support instead of a comprehensive framework for mobile development.

Installation
------------

```bash
go get -u github.com/maja42/gl/...
GOARCH=js go get -u -d github.com/maja42/gl/...
```

Usage
-----

This OpenGL binding has a ContextWatcher, which implements [glfw.ContextWatcher](https://godoc.org/github.com/maja/glfw#ContextWatcher)
interface. Recommended usage is with github.com/maja/glfw package, which accepts a ContextWatcher in its Init, and takes on the responsibility
of notifying it when context is made current or detached.

```Go
if err := glfw.Init(gl.ContextWatcher); err != nil {
	// Handle error.
}
defer glfw.Terminate()
```

If you're not using a ContextWatcher-aware glfw library, you must call methods of gl.ContextWatcher yourself whenever
you make a context current or detached.

```Go
window.MakeContextCurrent()
gl.ContextWatcher.OnMakeCurrent(nil)

glfw.DetachCurrentContext()
gl.ContextWatcher.OnDetach()
```

Note that if you are using a different windowing library than http://github.com/maja/glfw, all calls must happen within the same render thread.

```Go
renderThread := render.New()
defer renderThread.Terminate()

gl.Init(renderThread)
renderThread.Enqueue(true, func() {
    glfw.Init(gl.ContextWatcher)
    window, _ = glfw.CreateWindow(...)
})
```
