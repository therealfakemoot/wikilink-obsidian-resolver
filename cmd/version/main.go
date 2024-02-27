package main

import (
	"fmt"

	"github.com/therealfakemoot/wikilink-obsidian-resolver"
)

func main() {
	fmt.Println("Version:", resolver.Version)
	fmt.Println("Build:", resolver.Build)
}
