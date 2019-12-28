// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js

package gl

import "github.com/gopherjs/gopherjs/js"

type Enum int

type Attrib struct {
	Value int
}

type Program struct {
	*js.Object
}

type Shader struct {
	*js.Object
}

type Buffer struct {
	*js.Object
}

type Framebuffer struct {
	*js.Object
}

type Renderbuffer struct {
	*js.Object
}

type Texture struct {
	*js.Object
}

type Uniform struct {
	*js.Object
}

func (v Attrib) Valid() bool       { return v.Value != 0 }
func (v Program) Valid() bool      { return v.Object != nil }
func (v Shader) Valid() bool       { return v.Object != nil }
func (v Buffer) Valid() bool       { return v.Object != nil }
func (v Framebuffer) Valid() bool  { return v.Object != nil }
func (v Renderbuffer) Valid() bool { return v.Object != nil }
func (v Texture) Valid() bool      { return v.Object != nil }
func (v Uniform) Valid() bool      { return v.Object != nil }

func (v Attrib) String() string       { return fmt.Sprintf("Attrib(%d)", v.Value) }
func (v Program) String() string      { return fmt.Sprintf("Program(%d)", v.Value) }
func (v Shader) String() string       { return fmt.Sprintf("Shader(%d)", v.Value) }
func (v Buffer) String() string       { return fmt.Sprintf("Buffer(%d)", v.Value) }
func (v Framebuffer) String() string  { return fmt.Sprintf("Framebuffer(%d)", v.Value) }
func (v Renderbuffer) String() string { return fmt.Sprintf("Renderbuffer(%d)", v.Value) }
func (v Texture) String() string      { return fmt.Sprintf("Texture(%d)", v.Value) }
func (v Uniform) String() string      { return fmt.Sprintf("Uniform(%d)", v.Value) }
