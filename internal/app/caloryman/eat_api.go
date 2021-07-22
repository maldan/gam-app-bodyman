package caloryman

import (
	"fmt"
	"time"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type EatApi struct {
	File string
}

type EatApi_PostIndexArgs struct {
	Id        string
	ProductId string
	Amount    string
	Created   time.Time
}

// Get eat by id
func (r EatApi) GetIndex(args IdArgs) Eat {
	// Open file
	var list = DB_eatList()

	// Search
	for _, item := range list {
		if item.Id == args.Id {
			item.Product = DB_productById(item.ProductId)
			return item
		}
	}

	// Error
	restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
	return Eat{}
}

// Get full eat list
func (r EatApi) GetList() []Eat {
	return DB_eatList()
}

// Get filtered by date
func (r EatApi) GetFilterByDate(args DateArgs) []Eat {
	// List
	var list = DB_eatList()

	// Out result
	out := make([]Eat, 0)

	// Search
	for _, item := range list {
		// Not date, skip
		if !(item.Created.Year() == args.Date.Year() &&
			item.Created.Month() == args.Date.Month() &&
			item.Created.YearDay() == args.Date.YearDay()) {
			continue
		}

		// Get product for eat
		product := DB_productById(item.ProductId)
		item.Product = product

		// Calculate components
		item.Product.Protein = UCTo(UCFrom(item.Amount)/100*UCFrom(item.Product.Protein), "g")
		item.Product.Fat = UCTo(UCFrom(item.Amount)/100*UCFrom(item.Product.Fat), "g")
		item.Product.Carbohydrate = UCTo(UCFrom(item.Amount)/100*UCFrom(item.Product.Carbohydrate), "g")

		// Calculate calory
		item.Calory = UCFrom(item.Amount) / 100 * (UCFrom(product.Protein)*4 + UCFrom(product.Carbohydrate)*4 + UCFrom(product.Fat)*9)
		out = append(out, item)
	}

	return out
}

// Get total stat for date
func (f EatApi) GetTotalStatByDate(args DateArgs) map[string]float64 {
	var list = f.GetFilterByDate(DateArgs{Date: args.Date})
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
func (f EatApi) GetYearMap(args DateArgs) map[string]interface{} {
	out := map[string]interface{}{}
	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	start := time.Now()
	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp.TimeFormat(t2, "YYYY-MM-DD")] = f.GetTotalStatByDate(DateArgs{Date: t2})
	}
	fmt.Println(time.Since(start))

	return out
}

func (f EatApi) PostIndex(args EatApi_PostIndexArgs) {
	var eat []Eat
	docdb.Get(DataDir, "eat", &eat)
	eat = append(eat, Eat{
		Id:        xid.New().String(),
		ProductId: args.ProductId,
		Amount:    args.Amount,
		Created:   args.Created,
	})
	docdb.Save(DataDir, "eat", &eat)
}

func (f EatApi) PatchIndex(args EatApi_PostIndexArgs) {
	var eat []Eat
	docdb.Get(DataDir, "eat", &eat)
	item, itemId := cmhp.SliceFindR(eat, func(i interface{}) bool {
		return i.(Eat).Id == args.Id
	})

	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
	}
	eatItem := item.(Eat)
	eatItem.ProductId = args.ProductId
	eatItem.Amount = args.Amount
	eatItem.Created = args.Created
	eat[itemId] = eatItem

	docdb.Save(DataDir, "eat", &eat)
}

func (f EatApi) DeleteIndex(args DeleteIndexArgs) {
	var eat []Eat
	docdb.Get(DataDir, "eat", &eat)
	out := cmhp.SliceFilterR(eat, func(i interface{}) bool {
		return i.(Eat).Id != args.Id
	})
	docdb.Save(DataDir, "eat", &out)
}
