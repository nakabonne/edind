package edind

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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

// Open opens the file with the given editor
// TODO: エディターの選択肢増やせるようにする
// TODO: 出力先指定出来るようにする
func (e *Editor) Open(path string) error {
	// NOTE: bashの位置指定する場合
	//run := fmt.Sprintf("%s %s", editor.Name, Escape(path))
	//cmd := exec.Command(config.BashPath, "-c", run)

	//cmd := exec.Command(e.Name, path)
	cmd := exec.Command("open", "-t", "-W", path)
	fmt.Println(e.Name, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

var EDITORS = [][]string{
	{"open", "-t", "-W"}, // Opens with the default text editor on mac
	{"subl"},
	{"vim"},
	{"emacs"},
	{"mate", "-w"},
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
