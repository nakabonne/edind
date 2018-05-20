package edind

import (
	"io"
	"os/exec"
)

// Editor is an editor for opening files
type Editor struct {
	Name  string
	Flags []string

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// Open opens the file with the given editor
func (e *Editor) Open(path string) error {
	var cmd *exec.Cmd
	if len(e.Flags) >= 1 {
		e.Flags = append(e.Flags, path)
		cmd = exec.Command(e.Name, e.Flags...)
	} else {
		cmd = exec.Command(e.Name, path)
	}

	cmd.Stdin = e.Stdin
	cmd.Stdout = e.Stdout
	cmd.Stderr = e.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// SetStdin sets Standard input destination
func (e *Editor) SetStdin(r io.Reader) {
	e.Stdin = r
}

// SetStdout sets Standard output destination
func (e *Editor) SetStdout(w io.Writer) {
	e.Stdout = w
}

// SetStderr sets Standard error output destination
func (e *Editor) SetStderr(w io.Writer) {
	e.Stderr = w
}
