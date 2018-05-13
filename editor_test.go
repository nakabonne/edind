package edind

import "testing"

func TestOpen(t *testing.T) {
	editor := NewEditor()
	editor.AddDefaults("vi")
	err := editor.DetectEditor()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	err = editor.Open("editor.go")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
