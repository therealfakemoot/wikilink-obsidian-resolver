package resolver

import (
	// "errors"
	"io/fs"
	"log/slog"

	"go.abhg.dev/goldmark/wikilink"
)

var ()

func NewResolver(vaultRoot, blogRoot string) *Resolver {
	var r Resolver

	return &r
}

type Resolver struct {
	vaultFS fs.FS
	blogFS  fs.FS
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	out := make([]byte, 0)
	slog.Debug("entering ResolveWikilink",
		slog.Group("incoming node",
			slog.String("target", string(n.Target)),
			slog.String("fragment", string(n.Fragment)),
		),
	)

	return out, nil
}
