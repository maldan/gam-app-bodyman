package caloryman

import (
	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type WeightApi struct {
	Table string
}

func (f WeightApi) GetIndex(args IdArgs) (Weight, int) {
	// Find training
	list := f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return i.(Weight).Id == args.Id
	})

	// Not found
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Weight not found!")
	}

	return item.(Weight), itemId
}

func (f WeightApi) GetList() []Weight {
	var list []Weight
	docdb.Get(DataDir, f.Table, &list)
	return list
}

func (f WeightApi) PostIndex(args Weight) {
	list := f.GetList()
	args.Id = xid.New().String()
	list = append(list, args)
	docdb.Save(DataDir, f.Table, &list)
}

func (f WeightApi) PatchIndex(args Weight) {
	_, itemId := f.GetIndex(IdArgs{Id: args.Id})
	list := f.GetList()
	list[itemId] = args
	docdb.Save(DataDir, f.Table, &list)
}

func (f WeightApi) DeleteIndex(args DeleteIndexArgs) {
	list := f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Weight).Id != args.Id
	})
	docdb.Save(DataDir, f.Table, &out)
}
