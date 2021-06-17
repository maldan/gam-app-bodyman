package caloryman

import (
	"github.com/maldan/go-docdb"
	"github.com/rs/xid"
)

type FoodApi int

type FoodApi_PostIndexArgs struct {
	Name string
}

type DeleteIndexArgs struct {
	Id string
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
		Id:   xid.New().String(),
		Name: args.Name,
	})
	docdb.Save("db", "food", &food)
}

func (f FoodApi) DeleteIndex(args DeleteIndexArgs) {
	var food []Food
	docdb.Get("db", "food", &food)
	out := make([]Food, 0)

	for _, v := range food {
		if v.Id == args.Id {
			continue
		}
		out = append(out, v)
	}

	docdb.Save("db", "food", &out)
}
