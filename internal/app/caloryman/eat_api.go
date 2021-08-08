package caloryman

import (
	"sort"
	"time"

	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type EatApi struct {
}

// Get eat by id
func (r EatApi) GetIndex(args ArgsId) Eat {
	// Get file with eat
	var item Eat
	err := cmhp_file.ReadJSON(DataDir+"/eat/item/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
	}
	// Append product
	r1 := ProductApi{}
	item.Product = r1.GetSafeIndex(ArgsId{Id: item.ProductId})
	return item
}

// Get filtered by date
func (r EatApi) GetFilterByDate(args ArgsDate) []Eat {
	// Get id list
	idList := make([]string, 0)
	cmhp_file.ReadJSON(DataDir+"/eat/stat/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &idList)

	// Out result
	out := make([]Eat, 0)

	// Product api
	r1 := ProductApi{}

	// Search
	for _, id := range idList {
		// Get item or skip
		var item Eat
		err := cmhp_file.ReadJSON(DataDir+"/eat/item/"+id+".json", &item)
		if err != nil {
			continue
		}

		// Get product for eat
		product := r1.GetSafeIndex(ArgsId{Id: item.ProductId})
		item.Product = product

		// Calculate components
		item.Product.Protein = UCTo(UCFrom(item.Amount)/100*UCFrom(item.Product.Protein), "g")
		item.Product.Fat = UCTo(UCFrom(item.Amount)/100*UCFrom(item.Product.Fat), "g")
		item.Product.Carbohydrate = UCTo(UCFrom(item.Amount)/100*UCFrom(item.Product.Carbohydrate), "g")

		// Calculate calory
		item.Calory = UCFrom(item.Amount) / 100 * (UCFrom(product.Protein)*4 + UCFrom(product.Carbohydrate)*4 + UCFrom(product.Fat)*9)
		out = append(out, item)
	}

	// Sort by date
	sort.SliceStable(out, func(i, j int) bool {
		return out[j].Created.Unix() > out[i].Created.Unix()
	})

	return out
}

// Get total stat by date
func (f EatApi) GetTotalStatByDate(args ArgsDate) map[string]float64 {
	var list = f.GetFilterByDate(ArgsDate{Date: args.Date})
	out := map[string]float64{
		"calory":       0,
		"protein":      0,
		"carbohydrate": 0,
		"fat":          0,
		"water":        0,
	}

	for i := 0; i < len(list); i++ {
		out["calory"] += list[i].Calory
		out["protein"] += UCFrom(list[i].Product.Protein)
		out["carbohydrate"] += UCFrom(list[i].Product.Carbohydrate)
		out["fat"] += UCFrom(list[i].Product.Fat)

		if list[i].Product.Name == "Вода" {
			out["water"] += UCFrom(list[i].Amount)
		}
	}

	return out
}

// Get year calory stat
func (f EatApi) GetYearMap(args ArgsDate) map[string]interface{} {
	out := map[string]interface{}{}
	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp_time.Format(t2, "YYYY-MM-DD")] = f.GetTotalStatByDate(ArgsDate{Date: t2})
	}

	return out
}

// Add new eat
func (f EatApi) PostIndex(args Eat) {
	// Save to file
	args.Id = cmhp_crypto.UID(10)
	args.Amount = UCTo(UCFrom(args.Amount), "g")
	cmhp_file.WriteJSON(DataDir+"/eat/item/"+args.Id+".json", &args)

	// Get list
	idList := make([]string, 0)
	cmhp_file.ReadJSON(DataDir+"/eat/stat/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &idList)
	idList = append(idList, args.Id)
	cmhp_file.WriteJSON(DataDir+"/eat/stat/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &idList)
}

// Update eat
func (f EatApi) PatchIndex(args Eat) {
	args.Amount = UCTo(UCFrom(args.Amount), "g")
	cmhp_file.WriteJSON(DataDir+"/eat/item/"+args.Id+".json", &args)
}

// Delete eat info
func (r EatApi) DeleteIndex(args ArgsId) {
	item := r.GetIndex(args)
	cmhp_file.Delete(DataDir + "/eat/item/" + item.Id + ".json")

	// Get list
	idList := make([]string, 0)
	idListNew := make([]string, 0)
	cmhp_file.ReadJSON(DataDir+"/eat/stat/"+cmhp_time.Format(item.Created, "YYYY-MM-DD")+".json", &idList)

	// Remove id
	for _, id := range idList {
		if args.Id == id {
			continue
		}
		idListNew = append(idListNew, id)
	}

	// Save
	cmhp_file.WriteJSON(DataDir+"/eat/stat/"+cmhp_time.Format(item.Created, "YYYY-MM-DD")+".json", &idListNew)
}
