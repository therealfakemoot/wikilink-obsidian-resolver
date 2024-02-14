package resolver

import (
	"log"

	"go.abhg.dev/goldmark/wikilink"
)

type Resolver struct {
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	out := make([]byte, 0)
	log.Printf("Resolving node: %#+v\n", n)

	return out, nil
}
