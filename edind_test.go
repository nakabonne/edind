package edind

import "testing"

func TestExampleSuccess(t *testing.T) {
	editor, err := GetEditor()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	Open(editor, "editor.go")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
