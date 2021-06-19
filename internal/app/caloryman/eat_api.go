package caloryman

import (
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
}

func (f EatApi) GetIndex(args IdArgs) Eat {
	var eat []Eat
	docdb.Get("db", "eat", &eat)
	item, itemId := cmhp.SliceFindR(eat, func(i interface{}) bool {
		return i.(Eat).Id == args.Id
	})
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
	}
	return item.(Eat)
}

func (f EatApi) GetList() []Eat {
	var eat []Eat
	docdb.Get("db", "eat", &eat)
	return eat
}

func (f EatApi) PostIndex(args EatApi_PostIndexArgs) {
	var eat []Eat
	docdb.Get("db", "eat", &eat)
	eat = append(eat, Eat{
		Id:        xid.New().String(),
		ProductId: args.ProductId,
		Amount:    args.Amount,
		Created:   time.Now(),
	})
	docdb.Save("db", "eat", &eat)
}

func (f EatApi) PatchIndex(args EatApi_PostIndexArgs) {
	var eat []Eat
	docdb.Get("db", "eat", &eat)
	item, itemId := cmhp.SliceFindR(eat, func(i interface{}) bool {
		return i.(Eat).Id == args.Id
	})

	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Eat not found!")
	}
	eatItem := item.(Eat)
	eatItem.ProductId = args.ProductId
	eatItem.Amount = args.Amount
	eat[itemId] = eatItem

	docdb.Save("db", "eat", &eat)
}

func (f EatApi) DeleteIndex(args DeleteIndexArgs) {
	var eat []Eat
	docdb.Get("db", "eat", &eat)
	out := cmhp.SliceFilterR(eat, func(i interface{}) bool {
		return i.(Eat).Id != args.Id
	})
	docdb.Save("db", "eat", &out)
}
