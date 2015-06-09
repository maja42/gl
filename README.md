# gl [![Build Status](https://travis-ci.org/goxjs/gl.svg?branch=master)](https://travis-ci.org/goxjs/gl) [![GoDoc](https://godoc.org/github.com/goxjs/gl?status.svg)](https://godoc.org/github.com/goxjs/gl)

Package gl is a Go cross-platform binding for OpenGL, with an OpenGL ES 2-like API.

It supports:

- **OS X**, **Linux** and **Windows** via OpenGL 2.1 backend,

- **iOS** and **Android** via OpenGL ES 2.0 backend,

- **Modern Browsers** (desktop and mobile) via WebGL 1.0 backend.

This is a fork of golang.org/x/mobile/gl package with [CL 8793](https://go-review.googlesource.com/8793)
merged in and Windows support added. This package is fully functional, but may eventually become superceded by
the new x/mobile/gl plan. It will exist and be fully supported until it can be safely replaced by a better package.

Installation
------------

```bash
go get -u github.com/goxjs/gl/...
go get -u -d -tags=js github.com/goxjs/gl/...
```

Usage
-----

This OpenGL binding has a ContextWatcher, which implements [glfw.ContextWatcher](https://godoc.org/github.com/goxjs/glfw#ContextWatcher)
interface. Recommended usage is with github.com/goxjs/glfw package, which accepts a ContextWatcher in its Init, and takes on the responsibility
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
