package caloryman

import (
	"sort"
	"time"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type TrainingApi struct {
	Table string
}

func (f TrainingApi) GetIndex(args ArgsId) (Training, int) {
	// Find training
	list := f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return i.(Training).Id == args.Id
	})

	// Not found
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Training not found!")
	}

	e := ExerciseApi{Table: "exercise"}
	itemOut := item.(Training)
	itemOut.Exercise, _ = e.GetIndexSafe(ArgsId{Id: item.(Training).ExerciseId})

	return itemOut, itemId
}

func (f TrainingApi) GetList() []Training {
	var training []Training
	docdb.Get(DataDir, f.Table, &training)
	return training
}

func (f TrainingApi) GetFilterByDate(args ArgsDate) []interface{} {
	var list = f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Training).Created.Format("2006-01-02") == args.Date.Format("2006-01-02")
	})
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].(Training).Created.UTC().Unix() < out[j].(Training).Created.UTC().Unix()
	})

	e := ExerciseApi{Table: "exercise"}
	for itemId, item := range out {
		training := item.(Training)
		training.Exercise, _ = e.GetIndexSafe(ArgsId{Id: item.(Training).ExerciseId})
		out[itemId] = training
	}

	return out
}

func (f TrainingApi) GetTotalStatByDate(args ArgsDate) map[string]map[string]float64 {
	list := f.GetFilterByDate(ArgsDate{Date: args.Date})
	out := map[string]map[string]float64{
		"tool": {
			"own_weight":     0,
			"barbell":        0,
			"dumbbell":       0,
			"dumbbell_2":     0,
			"spring_gripper": 0,
			"total":          0,
		},
		"muscle": {
			"biceps":      0,
			"chest":       0,
			"palm":        0,
			"triceps":     0,
			"shrugs":      0,
			"quadriceps":  0,
			"buttock":     0,
			"front_delta": 0,
			"back_delta":  0,
			"abs":         0,
			"total":       0,
		},
	}

	for i := 0; i < len(list); i++ {
		if list[i].(Training).Exercise.Tool == "own_weight" {
			out["tool"][list[i].(Training).Exercise.Tool] += float64(list[i].(Training).Weight+60) * float64(list[i].(Training).Reps)
			out["tool"]["total"] += float64(list[i].(Training).Weight+60) * float64(list[i].(Training).Reps)

			for _, muscle := range list[i].(Training).Exercise.MuscleList {
				out["muscle"][muscle] += float64(list[i].(Training).Weight+60) * float64(list[i].(Training).Reps)
				out["muscle"]["total"] += float64(list[i].(Training).Weight+60) * float64(list[i].(Training).Reps)
			}
		} else {
			out["tool"][list[i].(Training).Exercise.Tool] += float64(list[i].(Training).Weight) * float64(list[i].(Training).Reps)
			out["tool"]["total"] += float64(list[i].(Training).Weight) * float64(list[i].(Training).Reps)

			for _, muscle := range list[i].(Training).Exercise.MuscleList {
				out["muscle"][muscle] += float64(list[i].(Training).Weight) * float64(list[i].(Training).Reps)
				out["muscle"]["total"] += float64(list[i].(Training).Weight) * float64(list[i].(Training).Reps)
			}
		}
	}

	return out
}

// Get year calory stat
func (f TrainingApi) GetYearMap(args ArgsDate) map[string]interface{} {
	out := map[string]interface{}{}

	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp.TimeFormat(t2, "YYYY-MM-DD")] = f.GetTotalStatByDate(ArgsDate{Date: t2})
	}

	return out
}

func (f TrainingApi) PostIndex(args Training) {
	list := f.GetList()
	args.Id = xid.New().String()
	/*if args.MuscleList == nil {
		args.MuscleList = make([]string, 0)
	}*/
	list = append(list, args)
	docdb.Save(DataDir, f.Table, &list)
}

func (f TrainingApi) PatchIndex(args Training) {
	/*if args.MuscleList == nil {
		args.MuscleList = make([]string, 0)
	}*/

	_, trainingId := f.GetIndex(ArgsId{Id: args.Id})
	list := f.GetList()
	list[trainingId] = args
	docdb.Save(DataDir, f.Table, &list)
}

func (f TrainingApi) DeleteIndex(args ArgsId) {
	list := f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Training).Id != args.Id
	})
	docdb.Save(DataDir, f.Table, &out)
}
