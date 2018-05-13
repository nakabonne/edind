package edind

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Editor is an editor for opening files
type Editor struct {
	Name  string
	Flags []string

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// DefaultEditors are detected editors
var DefaultEditors = []Editor{
	Editor{Name: "vim"},
	Editor{Name: "emacs"},
	Editor{Name: "nano"},
	Editor{Name: "subl"},
	Editor{Name: "atom"},
	Editor{Name: "open", Flags: []string{"-t", "-W"}},
	Editor{Name: "mate", Flags: []string{"-w"}},
}

// NewEditor returns an Editor
func NewEditor() *Editor {
	editor := &Editor{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return editor
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

	fmt.Println(e.Name, e.Flags, path)

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

// AddDefaults adds Editor to DefaultEditors
func (e *Editor) AddDefaults(name string, flags ...string) {
	DefaultEditors = append(DefaultEditors, Editor{Name: name, Flags: flags})
	return
}

// DetectEditor detects executable editor commands from the given PATH
func (e *Editor) DetectEditor() error {
	env := GetEnv()
	if name := env["EDITOR"]; name != "" {
		e.Name = name
		return nil
	}

	pathEnv := env["PATH"]
	for _, d := range DefaultEditors {
		if _, err := lookPath(d.Name, pathEnv); err == nil {
			e.Name = d.Name
			e.Flags = d.Flags
			return nil
		}
	}
	return fmt.Errorf("Could not find a default editor in the PATH")
}

func lookPath(file string, pathenv string) (string, error) {
	if strings.Contains(file, "/") {
		err := findExecutable(file)
		if err == nil {
			return file, nil
		}
		return "", err
	}
	if pathenv == "" {
		return "", errNotFound
	}
	for _, dir := range strings.Split(pathenv, ":") {
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}
		path := dir + "/" + file
		if err := findExecutable(path); err == nil {
			return path, nil
		}
	}
	return "", errNotFound
}

// ErrNotFound is the error resulting if a path search failed to find an executable file.
var errNotFound = errors.New("executable file not found in $PATH")

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}
