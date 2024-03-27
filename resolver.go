package resolver

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/therealfakemoot/go-obsidian"
	"go.abhg.dev/goldmark/wikilink"
)

var ErrNameNotResolved = errors.New("name could not be resolved")

func NewResolver(vaultRoot string, opts Opts) (*Resolver, error) {
	var r Resolver

	wd, err := os.Getwd()
	if err != nil {
		return &r, fmt.Errorf("couldn't get current working directory before opening vault dir: %w", err)
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

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	target := string(n.Target)

	for name, note := range r.Vault.Notes {
		if strings.Contains(name, target) {
			// FIXME: Critical: pull the URL format string from Hugo's config
			notePath := fmt.Sprintf("/%d/%02d/%s", note.Date.Year(), note.Date.Month(), name)

			return []byte(notePath), nil
		}
	}

	r.Log.Debug("name not resolved")

	return nil, ErrNameNotResolved
}
