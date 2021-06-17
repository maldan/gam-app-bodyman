package caloryman

import (
	"time"

	"github.com/maldan/go-docdb"
	"github.com/rs/xid"
)

type EatApi int

type EatApi_PostIndexArgs struct {
	ProductId string
	Amount    string
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

func (f EatApi) DeleteIndex(args DeleteIndexArgs) {
	var eat []Eat
	docdb.Get("db", "eat", &eat)
	out := make([]Eat, 0)

	for _, v := range eat {
		if v.Id == args.Id {
			continue
		}
		out = append(out, v)
	}

	docdb.Save("db", "eat", &out)
}
