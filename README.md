# edind

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/nakabonne/edind)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/nakabonne/edind/master/LICENSE)

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
