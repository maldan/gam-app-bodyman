package api

import (
	"strings"
	"time"

	"github.com/maldan/gam-app-caloryman/internal/app/caloryman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type WeightApi struct {
}

// Get weight
func (r WeightApi) GetIndex(args ArgsDate) core.Weight {
	// Get file with weight
	var item core.Weight
	err := cmhp_file.ReadJSON(core.DataDir+"/weight/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Weight not found!")
	}
	return item
}

// Get product list
func (r WeightApi) GetList() []core.Weight {
	files, _ := cmhp_file.List(core.DataDir + "/weight")
	out := make([]core.Weight, 0)
	for _, file := range files {
		date, _ := time.Parse("2006-01-02", strings.Replace(file.Name(), ".json", "", 1))
		out = append(out, r.GetIndex(ArgsDate{Date: date}))
	}
	return out
}

// Add new weight
func (r WeightApi) PostIndex(args core.Weight) {
	cmhp_file.WriteJSON(core.DataDir+"/weight/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &args)
}

// Update product
func (r WeightApi) PatchIndex(args core.Weight) {
	cmhp_file.WriteJSON(core.DataDir+"/weight/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &args)
}

// Delete product
func (r WeightApi) DeleteIndex(args ArgsDate) {
	cmhp_file.Delete(core.DataDir + "/weight/" + cmhp_time.Format(args.Date, "YYYY-MM-DD") + ".json")
}
