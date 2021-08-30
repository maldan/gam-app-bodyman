package caloryman

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/maldan/go-restserver"
)

var DataDir = "db"

func Start(frontFs embed.FS) {
	// Flags
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")
	var _ = flag.Bool("gui", false, "Use Gui")
	var initDev = flag.Bool("initDev", false, "Install dev")
	var _ = flag.Int("width", 1100, "Window Width")
	var _ = flag.Int("height", 900, "Window Height")
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()
	DataDir = *dataDir

	// Copy as dev app
	if *initDev {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		GamAppDir := strings.ReplaceAll(dirname, "\\", "/") + "/.gam-app"
		source, _ := os.Open(os.Args[0])
		os.MkdirAll(GamAppDir+"/dev-bodyman-v0.0.0", 0755)
		destination, err := os.Create(GamAppDir + "/dev-bodyman-v0.0.0/app.exe")
		if err != nil {
			panic(err)
		}
		io.Copy(destination, source)
	}

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
	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
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
	})
}
