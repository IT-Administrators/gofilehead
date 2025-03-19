package gofilehead

import (
	"io/fs"
	"log"
	"os"
	"path"
	"testing"
)

// Read specified file to byte array.
func readFile(path string) []byte {
	cont, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return cont
}

// List all files in specified directory.
func listFiles(dir string) []string {
	root := os.DirFS(dir)
	// Get only files ending with "".json".
	mdFiles, err := fs.Glob(root, "*.*")

	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, v := range mdFiles {
		files = append(files, path.Join(dir, v))
	}
	return files
}

func TestMatchFileHead(t *testing.T) {
	cont := readFile("./examples/Gopher.png")
	if !MatchFileHead(cont, "89504E470D0A1A0A") {
		t.Error("unknown filetype")
	}
}
