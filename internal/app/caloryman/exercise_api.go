package caloryman

import (
	"strings"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type ExerciseApi struct {
	Table string
}

func (f ExerciseApi) GetIndexSafe(args IdArgs) (Exercise, int) {
	// Find training
	list := f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return i.(Exercise).Id == args.Id
	})

	if itemId == -1 {
		return Exercise{}, -1
	}

	return item.(Exercise), itemId
}

func (f ExerciseApi) GetIndex(args IdArgs) (Exercise, int) {
	// Find training
	list := f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return i.(Exercise).Id == args.Id
	})

	// Not found
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Exercise not found!")
	}

	return item.(Exercise), itemId
}

func (f ExerciseApi) GetList() []Exercise {
	var list []Exercise
	docdb.Get(DataDir, f.Table, &list)
	return list
}

func (f ExerciseApi) GetByName(args NameArgs) Exercise {
	var list = f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return strings.Contains(strings.ToLower(i.(Exercise).Name), strings.ToLower(args.Name))
	})
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Exercise not found!")
	}
	return item.(Exercise)
}

func (f ExerciseApi) PostIndex(args Exercise) {
	list := f.GetList()
	args.Id = xid.New().String()
	if args.MuscleList == nil {
		args.MuscleList = make([]string, 0)
	}
	list = append(list, args)
	docdb.Save(DataDir, f.Table, &list)
}

func (f ExerciseApi) PatchIndex(args Exercise) {
	if args.MuscleList == nil {
		args.MuscleList = make([]string, 0)
	}

	_, trainingId := f.GetIndex(IdArgs{Id: args.Id})
	list := f.GetList()
	list[trainingId] = args
	docdb.Save(DataDir, f.Table, &list)
}

func (f ExerciseApi) DeleteIndex(args DeleteIndexArgs) {
	list := f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Exercise).Id != args.Id
	})
	docdb.Save(DataDir, f.Table, &out)
}
