package caloryman

import (
	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type FoodApi int

type FoodApi_PostIndexArgs struct {
	Id           string
	Name         string
	Fat          string
	Protein      string
	Carbohydrate string
}

type DeleteIndexArgs struct {
	Id string
}

func (f FoodApi) GetIndex(args IdArgs) Food {
	var food []Food
	docdb.Get("db", "food", &food)
	item, itemId := cmhp.SliceFindR(food, func(i interface{}) bool {
		return i.(Food).Id == args.Id
	})
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Food not found!")
	}
	return item.(Food)
}

func (f FoodApi) GetList() []Food {
	var food []Food
	docdb.Get("db", "food", &food)
	return food
}

func (f FoodApi) PostIndex(args FoodApi_PostIndexArgs) {
	var food []Food
	docdb.Get("db", "food", &food)
	food = append(food, Food{
		Id:              xid.New().String(),
		Name:            args.Name,
		Protein:         args.Protein,
		Fat:             args.Fat,
		Carbohydrate:    args.Carbohydrate,
		ComponentIdList: make([]string, 0),
	})
	docdb.Save("db", "food", &food)
}

func (f FoodApi) PatchIndex(args FoodApi_PostIndexArgs) {
	var food []Food
	docdb.Get("db", "food", &food)
	item, itemId := cmhp.SliceFindR(food, func(i interface{}) bool {
		return i.(Food).Id == args.Id
	})

	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Food not found!")
	}
	foodItem := item.(Food)
	foodItem.Name = args.Name
	food[itemId] = foodItem

	docdb.Save("db", "food", &food)
}

func (f FoodApi) DeleteIndex(args DeleteIndexArgs) {
	var food []Food
	docdb.Get("db", "food", &food)
	out := cmhp.SliceFilterR(food, func(i interface{}) bool {
		return i.(Food).Id != args.Id
	})
	docdb.Save("db", "food", &out)
}
