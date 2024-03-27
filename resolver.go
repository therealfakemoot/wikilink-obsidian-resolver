package resolver

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/therealfakemoot/go-obsidian"
	"go.abhg.dev/goldmark/wikilink"
)

var ErrNameNotResolved = errors.New("name could not be resolved")

func NewResolver(vaultRoot string, opts Opts) (*Resolver, error) {
	var r Resolver

	wd, err := os.Getwd()
	if err != nil {
		return &r, fmt.Errorf("couldn't get current working directory before opening vault dir: %s", err)
	}
	absPath := filepath.Join(wd, vaultRoot)
	v, err := obsidian.NewVault(absPath)
	if err != nil {
		return &r, fmt.Errorf("could not create resolver: %w", err)
	}

	r.Vault = v

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: opts.LogLevel,
	}))
	r.Log = l

	return &r, nil
}

type Opts struct {
	LogLevel slog.Level
}

type Resolver struct {
	Vault *obsidian.Vault
	Log   *slog.Logger
}

func (r *Resolver) Glob(target string) ([]string, error) {
	return nil, fmt.Errorf("Fart")
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	target := string(n.Target)

	matches, err := r.Glob(target)
	if err != nil {
		return nil, fmt.Errorf("could not glob for %q: %w", target, err)
	}

	r.Log.Debug("matches", slog.Any("files", matches))
	r.Log.Debug("matches count", slog.Int("len", len(matches)))

	if len(matches) > 0 {
		head, tail := filepath.Split(matches[0])
		ext := filepath.Ext(tail)

		if ext == ".md" {
			tail = tail[:len(tail)-len(ext)] + ".html"
		}

		return []byte(filepath.Join(head, tail)), nil
	}

	r.Log.Debug("name not resolved")

	return nil, ErrNameNotResolved
}
