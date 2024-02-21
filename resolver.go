package resolver

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

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

func (r *Resolver) Glob(target string) ([]string, error) {
	pattern := fmt.Sprintf("*%s*", target)

	l := r.Log.With(slog.String("pattern", pattern))
	l.Debug("searching for target")

	matches := make([]string, 0)

	walkFunc := func(path string, d fs.DirEntry, err error) error {
		l := l.With(slog.String("path", path), slog.String("dir", d.Name()))
		l.Debug("walkFunc stepping")

		if err != nil {
			return fmt.Errorf("walkFunc was handed an error: %w", err)
		}

		if strings.Contains(path, target) {
			matches = append(matches, path)
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
	target := string(n.Target)

	matches, err := r.Glob(target)
	if err != nil {
		return nil, fmt.Errorf("could not glob for %q: %w", target, err)
	}

	if len(matches) > 0 {
		head, tail := filepath.Split(matches[0])
		ext := filepath.Ext(tail)
		if ext == ".md" {
			tail = tail[:len(tail)-len(ext)] + ".html"
		}

		return []byte(filepath.Join(head, tail)), nil
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
