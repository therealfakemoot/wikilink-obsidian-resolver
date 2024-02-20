package resolver

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"go.abhg.dev/goldmark/wikilink"
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

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	out := make([]byte, 0)

	wildcard_glob := fmt.Sprintf("*%s*", n.Target)
	literal_glob, err := fs.Glob(r.vaultFS, wildcard_glob)
	log.Printf("%#+v\n", literal_glob)
	if err != nil {
		return nil, fmt.Errorf("could not locate target in provided vault: %w", err)
	}

	if len(literal_glob) > 0 {
		log.Println("match found")
		return []byte(literal_glob[0]), nil
	}

	return out, nil
}
