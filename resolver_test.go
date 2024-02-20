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
		expected string
	}{
		{
			name: "basic resolution",
			in: &wikilink.Node{
				Target: []byte("zk_topic_a"),
			},
			expected: "/Resources/blog/published/zk_topic_a",
		},
	}

	r, err := resolver.NewResolver("testdata/test_vault/")
	if err != nil {
		assert.Error(t, err, "could not load test vault")
	}

	for _, c := range cases {
		t.Parallel()
		t.Run(c.name, func(t *testing.T) {
			actual, err := r.ResolveWikilink(c.in)
			if err != nil {
				assert.Failf(t, "error resolving wikilink", "%s", err)
				// t.Logf("error resolving wikilink: %s\n", err)
				// t.Fail()
			}
			assert.Equal(t, c.expected, string(actual))
		})
	}
}
