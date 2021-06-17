module github.com/maldan/gam-app-caloryman

go 1.16

replace github.com/maldan/go-restserver => ../../../go_lib/restserver

replace github.com/maldan/go-docdb => ../../../go_lib/docdb

require (
	github.com/maldan/go-docdb v1.0.0
	github.com/maldan/go-restserver v1.1.2
	github.com/rs/xid v1.3.0 // indirect
	github.com/zserge/lorca v0.1.10
)
