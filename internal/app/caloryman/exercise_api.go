package caloryman

import (
	"strings"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-restserver"
)

type ExerciseApi struct {
}

// Get exercise by id without error
func (r ExerciseApi) GetSafeIndex(args ArgsId) Exercise {
	var item Exercise
	cmhp.FileReadAsJSON(DataDir+"/exercise/"+args.Id+".json", &item)
	return item
}

// Get exercise by id
func (r ExerciseApi) GetIndex(args ArgsId) Exercise {
	var item Exercise
	err := cmhp.FileReadAsJSON(DataDir+"/exercise/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Exercise not found!")
	}
	return item
}

// Get list of all exercises
func (r ExerciseApi) GetList() []Exercise {
	files, _ := cmhp.FileList(DataDir + "/exercise")
	out := make([]Exercise, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	return out
}

// Get exercise by name
func (f ExerciseApi) GetByName(args ArgsName) Exercise {
	var list = f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return strings.Contains(strings.ToLower(i.(Exercise).Name), strings.ToLower(args.Name))
	})
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Exercise not found!")
	}
	return item.(Exercise)
}

// Add new exercise
func (f ExerciseApi) PostIndex(args Exercise) {
	args.Id = cmhp.UID(10)
	cmhp.FileWriteAsJSON(DataDir+"/exercise/"+args.Id+".json", &args)
}

// Update exercise
func (f ExerciseApi) PatchIndex(args Exercise) {
	cmhp.FileWriteAsJSON(DataDir+"/exercise/"+args.Id+".json", &args)
}

// Delete exercise
func (f ExerciseApi) DeleteIndex(args ArgsId) {
	cmhp.FileDelete(DataDir + "/exercise/" + args.Id + ".json")
}
