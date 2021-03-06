package gas

import (
	"testing"
)

func TestGlobalOpen(t *testing.T) {
	file, err := Open("github.com/andrebq/gas/fs.go")
	if err != nil {
		t.Fatalf("Unable to open fs.go file. %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Error while closing the file. %v", err)
	}

	Refresh()

	file, err = Open("github.com/andrebq/gas/fs.go")
	if err != nil {
		t.Fatalf("Unable to open fs.go file. %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Error while closing the file. %v", err)
	}

	file, err = Open("fs.go")
	if err != nil {
		t.Fatalf("Unable to open fs.go file. %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Error while closing the file. %v", err)
	}

	Refresh()

	file, err = Open("fs.go")
	if err != nil {
		t.Fatalf("Unable to open fs.go file. %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Error while closing the file. %v", err)
	}
}
