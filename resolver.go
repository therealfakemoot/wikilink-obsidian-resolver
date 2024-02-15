package resolver

import (
	// "errors"
	"log/slog"

	"go.abhg.dev/goldmark/wikilink"
)

var ()

type Resolver struct {
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
