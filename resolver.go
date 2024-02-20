package resolver

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"go.abhg.dev/goldmark/wikilink"
)

var (
	ErrNameNotResolved = errors.New("name could not be resolved")
)

func NewResolver(vaultRoot string) (*Resolver, error) {
	var r Resolver

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("couldn't get current working directory before opening vault dir: %w", err)
	}
	absPath := filepath.Join(wd, vaultRoot)
	vaultFS := os.DirFS(absPath)

	r.vaultFS = vaultFS

	return &r, nil
}

type Resolver struct {
	vaultFS fs.FS
}

func (r *Resolver) Glob(pattern string) []string {
	files, err := fs.Glob(r.vaultFS, pattern)
	if err != nil {
		panic(err)
	}
	return files
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	out := make([]byte, 0)

	wildcard_glob := fmt.Sprintf("*%s*", n.Target)
	log.Printf("matching glob: %q\n", wildcard_glob)
	literal_glob, err := fs.Glob(r.vaultFS, wildcard_glob)
	if err != nil {
		return nil, fmt.Errorf("could not use glob to search: %w", err)
	}

	if len(literal_glob) > 0 {
		log.Println("match found")
		return []byte(literal_glob[0]), nil
	}

	return out, ErrNameNotResolved
}

func (r *Resolver) DebugFS() []string {
	files := make([]string, 0)

	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	}

	fs.WalkDir(r.vaultFS, ".", walkFunc)

	return files
}
