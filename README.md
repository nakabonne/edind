# edind

edind is library for opening files with editor

# Usage

```go
import "github.com/nakabonne/edind"

func main(){
	edind.AddEditors(
		[]string{"vi"},
		[]string{"oni", "-w"},
	)

	editor, _ := edind.DetectEditor()

	_ = editor.Open("editor.go")
}
```

# Installation

```
$ go get -u github.com/nakabonne/edind
```
