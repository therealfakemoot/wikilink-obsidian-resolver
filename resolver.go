package resolver

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"go.abhg.dev/goldmark/wikilink"
)

var (
	ErrNameNotResolved = errors.New("name could not be resolved")
)

func NewResolver(vaultRoot string, opts Opts) (*Resolver, error) {
	var r Resolver

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: opts.LogLevel,
	}))
	r.Log = l

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("couldn't get current working directory before opening vault dir: %w", err)
	}

	absPath := filepath.Join(wd, vaultRoot)
	vaultFS := os.DirFS(absPath)

	r.vaultFS = vaultFS

	return &r, nil
}

type Opts struct {
	LogLevel slog.Level
}

type Resolver struct {
	vaultFS fs.FS
	Log     *slog.Logger
}

func (r *Resolver) Glob(pattern string) ([]string, error) {
	matches := make([]string, 0)
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files, err := fs.Glob(r.vaultFS, pattern)
			if err != nil {
				return fmt.Errorf("error searching for files in dir %q: %w", d.Name(), err)
			}

			matches = append(matches, files...)
		}

		return nil
	}

	err := fs.WalkDir(r.vaultFS, ".", walkFunc)
	if err != nil {
		return matches, fmt.Errorf("error globbing for %q: %w", pattern, err)
	}

	return matches, nil
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	wildcardGlob := fmt.Sprintf("*%s*", string(n.Target))

	matches, err := r.Glob(wildcardGlob)
	if err != nil {
		return nil, fmt.Errorf("could not glob for %q: %w", wildcardGlob, err)
	}

	if len(matches) > 0 {
		return []byte(matches[0]), nil
	}

	return nil, ErrNameNotResolved
}

func (r *Resolver) DebugFS() ([]string, error) {
	files := make([]string, 0)

	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	}

	err := fs.WalkDir(r.vaultFS, ".", walkFunc)
	if err != nil {
		return files, fmt.Errorf("error walking FS: %w", err)
	}

	return files, nil
}
