# edind

edind is library for opening files with editor

# Usage

```go
import "github.com/nakabonne/edind"

func main(){
  editor, _ := GetEditor()
	_ = Open(editor, "editor.go")
}
```

# Installation

```
$ go get -u https://github.com/nakabonne/edind
```
