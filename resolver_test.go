package resolver_test

import (
	"errors"
	"testing"

	"github.com/therealfakemoot/wikilink-obsidian-resolver"

	"github.com/stretchr/testify/assert"

	"go.abhg.dev/goldmark/wikilink"
)

func Test_GoodWikilinkResolution(t *testing.T) {
	cases := []struct {
		name     string
		in       *wikilink.Node
		expected string
	}{
		{
			name: "target resolution",
			in: &wikilink.Node{
				Target: []byte("zk_topic_a"),
			},
			expected: "/2024/02/zk_topic_a",
		},
		{
			name: "target resolution b",
			in: &wikilink.Node{
				Target: []byte("zk_topic_b"),
			},
			expected: "/2024/01/zk_topic_b",
		},
		{
			name: "fragment resolution b",
			in: &wikilink.Node{
				Target:   []byte("zk_topic_b"),
				Fragment: []byte("zk topic about croagers"),
			},
			expected: "/2024/01/zk_topic_b",
		},
		{
			name: "fragment resolution a",
			in: &wikilink.Node{
				Target:   []byte("zk_topic_a"),
				Fragment: []byte("zk topic about glorbles"),
			},
			expected: "/2024/02/zk_topic_a",
		},
	}

	r, err := resolver.NewResolver("testdata/test_vault/", resolver.Opts{})
	if err != nil {
		t.Logf("couldn't create resolver: %s\n", err)
		t.Fail()
	}

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
