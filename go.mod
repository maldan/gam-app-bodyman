module github.com/maldan/gam-app-caloryman

go 1.16

// replace github.com/maldan/go-restserver => ../../../go_lib/restserver
// replace github.com/maldan/go-docdb => ../../../go_lib/docdb
// replace github.com/maldan/go-cmhp => ../../../go_lib/cmhp

require (
	github.com/maldan/go-cmhp v0.0.10
	github.com/maldan/go-restserver v1.2.7
	github.com/zserge/lorca v0.1.10
)
