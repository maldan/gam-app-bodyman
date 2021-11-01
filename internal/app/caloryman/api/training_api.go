package api

import (
	"sort"
	"time"

	"github.com/maldan/gam-app-caloryman/internal/app/caloryman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type TrainingApi struct {
}

// Get training by id
func (r TrainingApi) GetIndex(args ArgsId) core.Training {
	// Get file with eat
	var item core.Training
	err := cmhp_file.ReadJSON(core.DataDir+"/training/item/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Training not found!")
	}

	// Append exercise
	r1 := ExerciseApi{}
	item.Exercise = r1.GetIndex(ArgsId{Id: item.ExerciseId})
	return item
}

// Get training by date
func (r TrainingApi) GetFilterByDate(args ArgsDate) []core.Training {
	// Get id list
	idList := make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/training/stat/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &idList)

	// Out result
	out := make([]core.Training, 0)

	// Product api
	r1 := ExerciseApi{}

	// Search
	for _, id := range idList {
		// Get item or skip
		var item core.Training
		err := cmhp_file.ReadJSON(core.DataDir+"/training/item/"+id+".json", &item)
		if err != nil {
			continue
		}

		// Get product for eat
		item.Exercise = r1.GetSafeIndex(ArgsId{Id: item.ExerciseId})
		out = append(out, item)
	}

	// Sort by date
	sort.SliceStable(out, func(i, j int) bool {
		return out[j].Created.Unix() > out[i].Created.Unix()
	})

	return out
}

// Get total stat by date
func (r TrainingApi) GetTotalStatByDate(args ArgsDate) map[string]map[string]float64 {
	list := r.GetFilterByDate(ArgsDate{Date: args.Date})
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
		if list[i].Exercise.Tool == "own_weight" {
			out["tool"][list[i].Exercise.Tool] += float64(list[i].Weight+60) * float64(list[i].Reps)
			out["tool"]["total"] += float64(list[i].Weight+60) * float64(list[i].Reps)

			for _, muscle := range list[i].Exercise.MuscleList {
				out["muscle"][muscle] += float64(list[i].Weight+60) * float64(list[i].Reps)
				out["muscle"]["total"] += float64(list[i].Weight+60) * float64(list[i].Reps)
			}
		} else {
			out["tool"][list[i].Exercise.Tool] += float64(list[i].Weight) * float64(list[i].Reps)
			out["tool"]["total"] += float64(list[i].Weight) * float64(list[i].Reps)

			for _, muscle := range list[i].Exercise.MuscleList {
				out["muscle"][muscle] += float64(list[i].Weight) * float64(list[i].Reps)
				out["muscle"]["total"] += float64(list[i].Weight) * float64(list[i].Reps)
			}
		}
	}

	return out
}

// Get year training stat
func (r TrainingApi) GetYearMap(args ArgsDate) map[string]interface{} {
	out := map[string]interface{}{}

	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp_time.Format(t2, "YYYY-MM-DD")] = r.GetTotalStatByDate(ArgsDate{Date: t2})
	}

	return out
}

// Add new training
func (r TrainingApi) PostIndex(args core.Training) {
	// Save to file
	args.Id = cmhp_crypto.UID(10)
	cmhp_file.WriteJSON(core.DataDir+"/training/item/"+args.Id+".json", &args)

	// Get list
	idList := make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/training/stat/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &idList)
	idList = append(idList, args.Id)
	cmhp_file.WriteJSON(core.DataDir+"/training/stat/"+cmhp_time.Format(args.Created, "YYYY-MM-DD")+".json", &idList)
}

// Update training
func (r TrainingApi) PatchIndex(args core.Training) {
	cmhp_file.WriteJSON(core.DataDir+"/training/item/"+args.Id+".json", &args)
}

// Delete training
func (r TrainingApi) DeleteIndex(args ArgsId) {
	item := r.GetIndex(args)
	cmhp_file.Delete(core.DataDir + "/training/item/" + item.Id + ".json")

	// Get list
	idList := make([]string, 0)
	idListNew := make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/training/stat/"+cmhp_time.Format(item.Created, "YYYY-MM-DD")+".json", &idList)

	// Remove id
	for _, id := range idList {
		if args.Id == id {
			continue
		}
		idListNew = append(idListNew, id)
	}

	// Save
	cmhp_file.WriteJSON(core.DataDir+"/training/stat/"+cmhp_time.Format(item.Created, "YYYY-MM-DD")+".json", &idListNew)
}