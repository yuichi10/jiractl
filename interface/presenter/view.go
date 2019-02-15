package presenter

import "io"

type Line struct {
	Body      string
	Color     string
	Delimiter string
}
type Lines []*Line

type ViewFactory interface {
	CreateView(io.WriteCloser)
}

type Viewer interface {
	Show(Lines)
}
