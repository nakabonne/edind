# edind

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/nakabonne/edind)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/nakabonne/edind/master/LICENSE)

edind is library for opening files and directories with editor


## Get Started

### Usage

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

### Installation

```
$ go get -u github.com/nakabonne/edind
```

### Default Editors

By default the following editors are used.
If you want to add other editors to be detected, use [AddEditors](https://godoc.org/github.com/nakabonne/edind#AddEditors).

```
$ vim
$ emacs
$ nano
$ subl
$ atom
$ open -t -W
$ mate -w
```

### License

This library is licensed under the MIT License.
