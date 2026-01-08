package poker_test

import (
	poker "golang-learning/application"
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{file}
	tape.Write([]byte("abc"))
	file.Seek(0, io.SeekStart)
	newContent, _ := io.ReadAll(file)

	got := string(newContent)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
