# edind

edind is library for opening files with editor

# Usage

```go
import "github.com/nakabonne/edind"

func main(){
  editor, _ := edind.GetEditor()
	_ = edind.Open(editor, "editor.go")
}
```

# Installation

```
$ go get -u https://github.com/nakabonne/edind
```
