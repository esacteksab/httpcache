// How to use https://www.alexedwards.net/blog/how-to-manage-tool-dependencies-in-go-1.24-plus
module github.com/esacteksab/httpcache

go 1.25.5

tool (
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
	mvdan.cc/gofumpt
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20251219203646-944ab1f22d93 // indirect
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/telemetry v0.0.0-20251219171224-e0dd3dd4fbaf // indirect
	golang.org/x/tools v0.40.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.6.1 // indirect
	mvdan.cc/gofumpt v0.9.2 // indirect
)
