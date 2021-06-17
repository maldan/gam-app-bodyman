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

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "host")
	var port = flag.Int("port", 16000, "port")
	var gui = flag.Bool("gui", false, "gui")
	var width = flag.Int("width", 1280, "width")
	var height = flag.Int("height", 720, "height")
	flag.Parse()

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
			"food": new(FoodApi),
			"eat":  new(EatApi),
		},
	})
}
