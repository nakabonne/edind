package edind

import "testing"

func TestOpen(t *testing.T) {
	f := NewEditorFactory()
	f.AddChoices(
		[]string{"vi"},
		[]string{"oni", "-w"},
	)

	editor, err := f.DetectEditor()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	err = editor.Open("editor.go")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestAddChoices(t *testing.T) {
	f := NewEditorFactory()
	before := len(f.Choices)

	f.AddChoices([]string{"vi"})

	after := len(f.Choices)

	if after != before+1 {
		t.Fatalf("failed test")
	}
	if f.Choices[len(f.Choices)-1].Name != "vi" {
		t.Fatalf("failed test")
	}
}
