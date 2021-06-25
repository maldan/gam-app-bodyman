package caloryman

import (
	"sort"
	"time"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type EatApi int

type EatApi_PostIndexArgs struct {
	Id        string
	ProductId string
	Amount    string
	Created   time.Time
}

func (f EatApi) GetIndex(args IdArgs) Eat {
	var eat []Eat
	docdb.Get(DataDir, "eat", &eat)
	item, itemId := cmhp.SliceFindR(eat, func(i interface{}) bool {
		return i.(Eat).Id == args.Id
	})

	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
	}

	var p = new(ProductApi)
	productItem, _ := p.GetIndex(IdArgs{Id: item.(Eat).ProductId})
	eatItem := item.(Eat)
	eatItem.Product = productItem

	return eatItem
}

func (f EatApi) GetList() []Eat {
	var eat []Eat
	docdb.Get(DataDir, "eat", &eat)
	return eat
}

func (f EatApi) GetFilterByDate(args DateArgs) []interface{} {
	var eatList = f.GetList()

	out := cmhp.SliceFilterR(eatList, func(i interface{}) bool {
		return i.(Eat).Created.Format("2006-01-02") == args.Date.Format("2006-01-02")
	})

	var p = new(ProductApi)

	for i := 0; i < len(out); i++ {
		productItem, _ := p.GetIndex(IdArgs{Id: out[i].(Eat).ProductId})
		eatItem := out[i].(Eat)
		eatItem.Product = productItem
		eatItem.Product.Protein = UCTo(UCFrom(eatItem.Amount)/100*UCFrom(eatItem.Product.Protein), "g")
		eatItem.Product.Fat = UCTo(UCFrom(eatItem.Amount)/100*UCFrom(eatItem.Product.Fat), "g")
		eatItem.Product.Carbohydrate = UCTo(UCFrom(eatItem.Amount)/100*UCFrom(eatItem.Product.Carbohydrate), "g")
		eatItem.Calory = UCFrom(eatItem.Amount) / 100 * (UCFrom(productItem.Protein)*4 + UCFrom(productItem.Carbohydrate)*4 + UCFrom(productItem.Fat)*9)
		out[i] = eatItem
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].(Eat).Created.UTC().Unix() < out[j].(Eat).Created.UTC().Unix()
	})

	return out
}

func (f EatApi) GetTotalStatByDate(args DateArgs) map[string]float64 {
	eatList := f.GetFilterByDate(DateArgs{Date: args.Date})
	out := map[string]float64{
		"calory":       0,
		"protein":      0,
		"carbohydrate": 0,
		"fat":          0,
		"water":        0,
	}

	for i := 0; i < len(eatList); i++ {
		out["calory"] += eatList[i].(Eat).Calory
		out["protein"] += UCFrom(eatList[i].(Eat).Product.Protein)
		out["carbohydrate"] += UCFrom(eatList[i].(Eat).Product.Carbohydrate)
		out["fat"] += UCFrom(eatList[i].(Eat).Product.Fat)

		if eatList[i].(Eat).Product.Name == "Вода" {
			out["water"] += UCFrom(eatList[i].(Eat).Amount)
		}
	}

	return out
}

// Get year calory stat
func (f EatApi) GetYearMap(args DateArgs) map[string]interface{} {
	out := map[string]interface{}{}

	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp.TimeFormat(t2, "YYYY-MM-DD")] = f.GetTotalStatByDate(DateArgs{Date: t2})
	}

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
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
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
