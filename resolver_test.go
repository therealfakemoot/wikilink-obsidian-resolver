package resolver_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/therealfakemoot/wikilink-obsidian-resolver"

	"github.com/stretchr/testify/assert"

	"go.abhg.dev/goldmark/wikilink"
)

func Test_BasicWikilinkResolution(t *testing.T) {
	t.Parallel()
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
			expected: "Resources/zk/zk_topic_a.html",
		},
		{
			name: "basic resolution b",
			in: &wikilink.Node{
				Target: []byte("zk_topic_b"),
			},
			expected: "Resources/zk/zk_topic_b.html",
		},
		{
			name: "basic resolution c",
			in: &wikilink.Node{
				Target: []byte("zk_topic_c"),
			},
			expected: "Resources/zk/zk_topic_c.html",
		},
	}

	r := resolver.NewResolver("testdata/test_vault/", resolver.Opts{LogLevel: slog.LevelDebug})
	// t.Logf("all files in vaultFS: %v\n", r.DebugFS())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := r.ResolveWikilink(c.in)
			if err != nil {
				if errors.Is(err, resolver.ErrNameNotResolved) {
					assert.Error(t, err, "name not correctly resolved")
				}
				assert.Errorf(t, err, "error resolving wikilink %q", c.in)
			}
			assert.Equal(t, c.expected, string(actual))
		})
	}
}
