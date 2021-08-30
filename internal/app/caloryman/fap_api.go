package caloryman

import (
	"strings"
	"time"

	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type FapApi struct {
}

// Get weight
func (r FapApi) GetIndex(args ArgsDate) Fap {
	// Get file with fap
	var item Fap
	err := cmhp_file.ReadJSON(DataDir+"/fap/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Fap not found!")
	}
	return item
}

// Get product list
func (r FapApi) GetList() []Fap {
	files, _ := cmhp_file.List(DataDir + "/fap")
	out := make([]Fap, 0)
	for _, file := range files {
		date, _ := time.Parse("2006-01-02", strings.Replace(file.Name(), ".json", "", 1))
		out = append(out, r.GetIndex(ArgsDate{Date: date}))
	}
	return out
}

// Add new fap
func (r FapApi) PostIndex(args Fap) {
	var item Fap
	cmhp_file.ReadJSON(DataDir+"/fap/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &item)
	item.Amount += args.Amount
	item.Created = args.Created
	cmhp_file.WriteJSON(DataDir+"/fap/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &item)
}

// Delete fap
func (r FapApi) DeleteIndex(args ArgsDate) {
	cmhp_file.Delete(DataDir + "/fap/" + cmhp_time.Format(args.Date, "YYYY-MM-DD") + ".json")
}