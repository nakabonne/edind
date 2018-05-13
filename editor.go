package edind

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Editor struct {
	Name string
}

func GetEditor() (editor Editor, err error) {
	env := GetEnv()
	name := env["EDITOR"]
	if name == "" {
		//logError("Could not find $EDITOR")
		name = detectEditor(env["PATH"])
		if name == "" {
			err = fmt.Errorf("Could not find a default editor in the PATH")
			return editor, err
		}
	}
	editor.Name = name
	return editor, nil
}

var EDITORS = [][]string{
	{"subl"},
	{"vim"},
	{"emacs"},
	{"mate", "-w"},
	{"open", "-t", "-W"}, // Opens with the default text editor on mac
	{"nano"},
}

func detectEditor(pathenv string) string {
	for _, editor := range EDITORS {
		if _, err := lookPath(editor[0], pathenv); err == nil {
			return strings.Join(editor, " ")
		}
	}
	return ""
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
