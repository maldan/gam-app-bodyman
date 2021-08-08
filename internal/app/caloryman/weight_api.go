package caloryman

import (
	"strings"
	"time"

	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type WeightApi struct {
}

// Get weight
func (r WeightApi) GetIndex(args ArgsDate) Weight {
	// Get file with weight
	var item Weight
	err := cmhp_file.ReadJSON(DataDir+"/weight/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Weight not found!")
	}
	return item
}

// Get product list
func (r WeightApi) GetList() []Weight {
	files, _ := cmhp_file.List(DataDir + "/weight")
	out := make([]Weight, 0)
	for _, file := range files {
		date, _ := time.Parse("2006-01-02", strings.Replace(file.Name(), ".json", "", 1))
		out = append(out, r.GetIndex(ArgsDate{Date: date}))
	}
	return out
}

// Add new weight
func (r WeightApi) PostIndex(args Weight) {
	cmhp_file.WriteJSON(DataDir+"/weight/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &args)
}

// Update product
func (r WeightApi) PatchIndex(args Weight) {
	cmhp_file.WriteJSON(DataDir+"/weight/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &args)
}

// Delete product
func (r WeightApi) DeleteIndex(args ArgsDate) {
	cmhp_file.Delete(DataDir + "/weight/" + cmhp_time.Format(args.Date, "YYYY-MM-DD") + ".json")
}
