package bzip2exe

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

type writer struct {
	w   io.Writer
	cmd *exec.Cmd
}

func NewWriter(out io.Writer) io.Writer {

	cmd := exec.Command("bzip2", "-c")

	cmd.Stdout = out
	cmd.Stderr = os.Stderr

	w := &writer{w: out, cmd: cmd}

	return w

}

func (w *writer) Write(data []byte) (int, error) {
	stdin := strings.NewReader(string(data))
	w.cmd.Stdin = stdin

	err := w.cmd.Run()

	return len(data), err

}
