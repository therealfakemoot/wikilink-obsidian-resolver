package main

import (
	"fmt"

	"github.com/therealfakemoot/wikilink-obsidian-resolver"
)

func main() {

	r, err := resolver.NewResolver("testdata/test_vault/")
	if err != nil {

	}

	fmt.Println(r.DebugFS())
	fmt.Printf("%#+v\n", r.Glob("*"))
}
