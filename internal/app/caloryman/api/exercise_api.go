package api

import (
	"strings"

	"github.com/maldan/gam-app-caloryman/internal/app/caloryman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"github.com/maldan/go-restserver"
)

type ExerciseApi struct {
}

// Get exercise by id without error
func (r ExerciseApi) GetSafeIndex(args ArgsId) core.Exercise {
	var item core.Exercise
	cmhp_file.ReadJSON(core.DataDir+"/exercise/"+args.Id+".json", &item)
	return item
}

// Get exercise by id
func (r ExerciseApi) GetIndex(args ArgsId) core.Exercise {
	var item core.Exercise
	err := cmhp_file.ReadJSON(core.DataDir+"/exercise/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Exercise not found!")
	}
	return item
}

// Get list of all exercises
func (r ExerciseApi) GetList() []core.Exercise {
	files, _ := cmhp_file.List(core.DataDir + "/exercise")
	out := make([]core.Exercise, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	return out
}

// Get exercise by name
func (f ExerciseApi) GetByName(args ArgsName) core.Exercise {
	var list = f.GetList()
	item, itemId := cmhp_slice.FindR(list, func(i interface{}) bool {
		return strings.Contains(strings.ToLower(i.(core.Exercise).Name), strings.ToLower(args.Name))
	})
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Exercise not found!")
	}
	return item.(core.Exercise)
}

// Add new exercise
func (f ExerciseApi) PostIndex(args core.Exercise) {
	args.Id = cmhp_crypto.UID(10)
	cmhp_file.WriteJSON(core.DataDir+"/exercise/"+args.Id+".json", &args)
}

// Update exercise
func (f ExerciseApi) PatchIndex(args core.Exercise) {
	cmhp_file.WriteJSON(core.DataDir+"/exercise/"+args.Id+".json", &args)
}

// Delete exercise
func (f ExerciseApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/exercise/" + args.Id + ".json")
}
