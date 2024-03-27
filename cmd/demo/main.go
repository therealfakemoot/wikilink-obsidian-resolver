package main

import (
	"log"

	"github.com/therealfakemoot/wikilink-obsidian-resolver"
)

var vaultRoot = "testdata/test_vault/"

func main() {
	r, err := resolver.NewResolver(vaultRoot, resolver.Opts{})
	if err != nil {
		log.Fatalf("couldn't build resolver: %s\n", err)
	}

	for _, note := range r.Vault.Notes {
		log.Printf("%#+v\n", note)
	}
}
