package edind

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// EditorFactory is a factory that creates Editors.
type EditorFactory struct {
	Choices []Editor
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

// NewEditorFactory returns a EditorFactory
func NewEditorFactory() (f *EditorFactory) {
	f = &EditorFactory{Choices: DefaultEditors}
	return
}

// DetectEditor detects executable editor commands from the given PATH
func (f *EditorFactory) DetectEditor() (*Editor, error) {
	env := GetEnv()
	editor := &Editor{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if name := env["EDITOR"]; name != "" {
		editor.Name = name
		return editor, nil
	}

	pathEnv := env["PATH"]
	for _, e := range f.Choices {
		if _, err := f.lookPath(e.Name, pathEnv); err == nil {
			editor.Name = e.Name
			editor.Flags = e.Flags
			return editor, nil
		}
	}
	return nil, fmt.Errorf("Could not find a default editor in the PATH")
}

// AddChoices adds choices to detect an executable editor
func (f *EditorFactory) AddChoices(editors ...[]string) {
	for _, e := range editors {
		if len(e) > 1 {
			f.Choices = append(f.Choices, Editor{Name: e[0], Flags: e[1:]})
		} else if len(e) == 1 {
			f.Choices = append(f.Choices, Editor{Name: e[0]})
		}
	}
	return
}

// errNotFound is the error resulting if a path search failed to find an executable file.
var errNotFound = errors.New("executable file not found in $PATH")

func (f *EditorFactory) lookPath(file string, pathenv string) (string, error) {
	if strings.Contains(file, "/") {
		err := f.findExecutable(file)
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
		if err := f.findExecutable(path); err == nil {
			return path, nil
		}
	}
	return "", errNotFound
}

func (f *EditorFactory) findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}
