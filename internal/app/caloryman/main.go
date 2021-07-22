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
	"github.com/zserge/lorca"
)

var DataDir = "db"

func Start(frontFs embed.FS) {
	// Flags
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")
	var gui = flag.Bool("gui", false, "Use Gui")
	var initDev = flag.Bool("initDev", false, "Install dev")
	var width = flag.Int("width", 1100, "Window Width")
	var height = flag.Int("height", 900, "Window Height")
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

	// GUI
	if *gui {
		go (func() {
			ui, _ := lorca.New("", "", *width, *height)
			defer ui.Close()
			ui.Load(fmt.Sprintf("http://%s:%d/", *host, *port))
			<-ui.Done()
			os.Exit(0)
		})()
	}

	/*d, _ := cmhp.FileReadAsBin("C:/Users/black/.gam-data/maldan-gam-app-caloryman/exercise.data")
	x, _ := cmhp.DataDecompress(d)
	cmhp.FileWriteAsBin("C:/Users/black/.gam-data/maldan-gam-app-caloryman/exercise.json", x)

	var list []Exercise
	cmhp.FileReadAsJSON(DataDir+"/exercise.json", &list)

	for _, item := range list {
		cmhp.FileWriteAsJSON(DataDir+"/exercise/"+item.Id+".json", &item)

		// fmt.Println(item)
		//itemIdList := make([]string, 0)
		//cmhp.FileReadAsJSON(DataDir+"/training/stat/"+cmhp.TimeFormat(item.Created, "YYYY-MM-DD")+".json", &itemIdList)
		// cmhp.FileWriteAsJSON(DataDir+"/eat/item/"+item.Id+".json", &item)
		//itemIdList = append(itemIdList, item.Id)
		//cmhp.FileWriteAsJSON(DataDir+"/training/stat/"+cmhp.TimeFormat(item.Created, "YYYY-MM-DD")+".json", &itemIdList)
	}*/

	/*d, _ := cmhp.FileReadAsBin("C:/Users/black/.gam-data/maldan-gam-app-caloryman/weight.data")
	x := cmhp.DataDecompress(d)
	cmhp.FileWriteAsBin("C:/Users/black/.gam-data/maldan-gam-app-caloryman/weight.json", x)*/

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
		},
	})
}
