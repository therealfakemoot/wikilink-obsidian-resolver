package resolver_test

import (
	"testing"

	"github.com/therealfakemoot/wikilink-obsidian-resolver"

	"github.com/stretchr/testify/assert"

	"go.abhg.dev/goldmark/wikilink"
)

func Test_BasicWikilinkResolution(t *testing.T) {
	cases := []struct {
		name     string
		in       *wikilink.Node
		expected []byte
	}{
		{
			name: "no alias, no fragment",
			in: &wikilink.Node{
				Target: []byte("target1"),
			},
		},
	}

	r := resolver.Resolver{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := r.ResolveWikilink(c.in)
			if err != nil {
				assert.Failf(t, "error resolving wikilink", "%s", err)
				// t.Logf("error resolving wikilink: %s\n", err)
				// t.Fail()
			}
			assert.Equal(t, actual, c.expected)
		})
	}
}
