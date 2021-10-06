package caloryman

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/gam-app-caloryman/internal/app/caloryman/api"
	"github.com/maldan/gam-app-caloryman/internal/app/caloryman/core"
	"github.com/maldan/go-rapi"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-rapi/rapi_rest"
	"github.com/maldan/go-rapi/rapi_vfs"
)

func Start(frontFs embed.FS) {
	// Flags
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.Int("clientPort", 8080, "Client Port")
	_ = flag.String("appId", "id", "App id")

	flag.Parse()

	core.DataDir = *dataDir

	/*m := make([]struct {
		Created time.Time `json:"created"`
	}, 0)
	cmhp_file.ReadJSON("./fap.json", &m)
	fff := FapApi{}
	for _, xx := range m {
		loc, _ := time.LoadLocation("Europe/Moscow")
		fff.PostIndex(Fap{
			Amount:  1,
			Created: xx.Created.In(loc),
		})
	}*/

	// Rest start
	/*restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"product":   ProductApi{},
			"eat":       EatApi{},
			"component": new(ComponentApi),
			"note":      new(NoteApi),
			"training":  TrainingApi{},
			"exercise":  ExerciseApi{},
			"weight":    WeightApi{},
			"fap":       FapApi{},
		},
	})*/

	// Start server
	rapi.Start(rapi.Config{
		Host: fmt.Sprintf("%s:%d", *host, *port),
		Router: map[string]rapi_core.Handler{
			"/": rapi_vfs.VFSHandler{
				Root: "frontend/build",
				Fs:   frontFs,
			},
			"/api": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"product":  api.ProductApi{},
					"eat":      api.EatApi{},
					"note":     api.NoteApi{},
					"training": api.TrainingApi{},
					"exercise": api.ExerciseApi{},
					"weight":   api.WeightApi{},
					"fap":      api.FapApi{},
				},
			},
		},
		DbPath: core.DataDir,
	})
}
