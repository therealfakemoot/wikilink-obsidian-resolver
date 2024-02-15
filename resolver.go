package resolver

import (
	// "errors"
	"log"

	"go.abhg.dev/goldmark/wikilink"
)

var ()

type Resolver struct {
}

func (r *Resolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	out := make([]byte, 0)
	log.Printf("%v\n", map[string]string{
		"Target":   string(n.Target),
		"Fragment": string(n.Fragment),
	})

	return out, nil
}
