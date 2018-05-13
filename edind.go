package edind

import (
	"os"
	"os/exec"
)

// Open opens the file with the given editor
// TODO: エディターの選択肢増やせるようにする
// TODO: 出力先指定出来るようにする
func Open(editor Editor, path string) error {
	// NOTE: bashの位置指定する場合
	//run := fmt.Sprintf("%s %s", editor.Name, Escape(path))
	//cmd := exec.Command(config.BashPath, "-c", run)

	cmd := exec.Command(editor.Name, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
