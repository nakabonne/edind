package edind

import "testing"

func TestOpen(t *testing.T) {
	editor := NewEditor()
	err := editor.DetectEditor()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	err = editor.Open("editor.go")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestAddDefaultEditors(t *testing.T) {
	before := len(DefaultEditors)

	editor := NewEditor()
	editor.AddDefaults("vi")

	after := len(DefaultEditors)

	if after != before+1 {
		t.Fatalf("failed test")
	}
	if DefaultEditors[len(DefaultEditors)-1].Name != "vi" {
		t.Fatalf("failed test")
	}
}
