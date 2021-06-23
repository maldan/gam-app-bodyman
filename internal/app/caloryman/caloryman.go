package caloryman

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/zserge/lorca"
)

var DataDir = "db"

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	var gui = flag.Bool("gui", false, "Use Gui")
	var width = flag.Int("width", 1100, "Window Width")
	var height = flag.Int("height", 900, "Window Height")
	var dataDir = flag.String("data-dir", "db", "Data Directory")
	_ = flag.String("app-id", "id", "App id")
	flag.Parse()
	DataDir = *dataDir

	if *gui {
		go (func() {
			ui, _ := lorca.New("", "", *width, *height)
			defer ui.Close()
			ui.Load(fmt.Sprintf("http://%s:%d/", *host, *port))
			<-ui.Done()
			os.Exit(0)
		})()
	}

	docdb.Start()

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"product":   new(ProductApi),
			"eat":       new(EatApi),
			"component": new(ComponentApi),
			"note":      new(NoteApi),
		},
	})
}
