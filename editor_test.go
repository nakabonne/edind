package edind

import "testing"

func TestExampleSuccess(t *testing.T) {
	editor, err := NewEditor()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	err = editor.Open("editor.go")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
