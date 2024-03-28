package resolver

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/therealfakemoot/go-obsidian"
	"go.abhg.dev/goldmark/wikilink"
	"go.uber.org/zap"
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

	r.Log = v.Logger

	return &r, nil
}

type Opts struct {
	LogLevel zap.AtomicLevel
}

type Resolver struct {
	Vault *obsidian.Vault
	Log   *zap.Logger
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	target := string(n.Target)

	for name, note := range r.Vault.Notes {
		if name == target {
			notePath := fmt.Sprintf("/%d/%02d/%s", note.Date.Year(), note.Date.Month(), name)

			return []byte(notePath), nil
		}
	}

	r.Log.Info("name not resolved")

	return nil, ErrNameNotResolved
}
