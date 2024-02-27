pushd .
cd ~/personal/obsidian-hugo/
CGO_ENABLED=1 go build -tags extended -o ~/personal/wikilink-obsidian-resolver/obs-hugo .
popd
./obs-hugo serve --config testdata/config.toml --contentDir testdata/test_vault --printPathWarnings --printUnusedTemplates --templateMetrics --templateMetricsHints
