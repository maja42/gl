// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux
// +build arm arm64

package gl

var ContextWatcher contextWatcher

type contextWatcher struct{}

func (contextWatcher) OnMakeCurrent(context interface{}) {
	// context must be a WebGLRenderingContext *js.Object.
	c = context.(*js.Object)
}
func (contextWatcher) OnDetach() {
	c = nil
}
