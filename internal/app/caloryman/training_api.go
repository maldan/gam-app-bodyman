package caloryman

import (
	"sort"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type TrainingApi struct {
	Table string
}

func (f TrainingApi) GetIndex(args IdArgs) (Training, int) {
	// Find training
	trainingList := f.GetList()
	item, itemId := cmhp.SliceFindR(trainingList, func(i interface{}) bool {
		return i.(Training).Id == args.Id
	})

	// Not found
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Training not found!")
	}

	return item.(Training), itemId
}

func (f TrainingApi) GetList() []Training {
	var training []Training
	docdb.Get(DataDir, f.Table, &training)
	return training
}

func (f TrainingApi) GetFilterByDate(args DateArgs) []interface{} {
	var list = f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Training).Created.Format("2006-01-02") == args.Date.Format("2006-01-02")
	})
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].(Training).Created.UTC().Unix() < out[j].(Training).Created.UTC().Unix()
	})
	return out
}

func (f TrainingApi) PostIndex(args Training) {
	list := f.GetList()
	args.Id = xid.New().String()
	if args.MuscleList == nil {
		args.MuscleList = make([]string, 0)
	}
	list = append(list, args)
	docdb.Save(DataDir, f.Table, &list)
}

func (f TrainingApi) PatchIndex(args Training) {
	if args.MuscleList == nil {
		args.MuscleList = make([]string, 0)
	}

	_, trainingId := f.GetIndex(IdArgs{Id: args.Id})
	list := f.GetList()
	list[trainingId] = args
	docdb.Save(DataDir, f.Table, &list)
}

func (f TrainingApi) DeleteIndex(args DeleteIndexArgs) {
	list := f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Training).Id != args.Id
	})
	docdb.Save(DataDir, f.Table, &out)
}
