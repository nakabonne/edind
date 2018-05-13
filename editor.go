package edind

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Editor is an editor for opening files
type Editor struct {
	Name  string
	Flags []string
}

// NewEditor returns an Editor
// TODO: エディターの選択肢増やせるようにする
// TODO: 出力先指定出来るようにする
func NewEditor() (editor *Editor, err error) {
	env := GetEnv()
	editor = &Editor{}
	if name := env["EDITOR"]; name != "" {
		editor.Name = name
		return editor, nil
	}

	//logError("Could not find $EDITOR")
	err = editor.DetectEditor(env["PATH"])
	if err != nil {
		return nil, err
	}
	return editor, nil
}

// Open opens the file with the given editor
func (e *Editor) Open(path string) error {
	// NOTE: bashの位置指定する場合
	//run := fmt.Sprintf("%s %s", editor.Name, Escape(path))
	//cmd := exec.Command(config.BashPath, "-c", run)

	var cmd *exec.Cmd
	if len(e.Flags) >= 1 {
		e.Flags = append(e.Flags, path)
		cmd = exec.Command(e.Name, e.Flags...)
	} else {
		cmd = exec.Command(e.Name, path)
	}

	fmt.Println(e.Name, e.Flags, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

var DefaultEditors = []Editor{
	Editor{Name: "vim"},
	Editor{Name: "emacs"},
	Editor{Name: "nano"},
	Editor{Name: "subl"},
	Editor{Name: "atom"},
	Editor{Name: "open", Flags: []string{"-t", "-W"}},
	Editor{Name: "mate", Flags: []string{"-w"}},
}

// DetectEditor detects executable editor commands from the given PATH
func (e *Editor) DetectEditor(pathenv string) error {
	for _, d := range DefaultEditors {
		if _, err := lookPath(d.Name, pathenv); err == nil {
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
