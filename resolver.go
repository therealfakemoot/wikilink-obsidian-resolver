package resolver

import (
	"go.abhg.dev/goldmark/wikilink"
)

type Resolver struct {
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error)
