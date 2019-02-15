package view

import (
	"fmt"
	"io"
	"os"

	"github.com/yuichi10/jiractl/interface/presenter"
)

type StdoutViewer struct {
	writer io.WriteCloser
}

func NewStdoutViewer() StdoutViewer {
	return StdoutViewer{writer: os.Stdout}
}

func (v StdoutViewer) Show(lines presenter.Lines) {
	for _, line := range lines {
		fmt.Fprintf(v.writer, "\x1b[%sm%s\x1b[0m", line.Color, line.Body)
		fmt.Fprint(v.writer, line.Delimiter)
	}
}
