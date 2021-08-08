package caloryman

import (
	"sort"
	"strings"
	"time"

	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-restserver"
)

type NoteApi int

type NoteApi_PostIndexArgs struct {
	Id          string
	Description string
	Created     time.Time
}

// Get note by id
func (r NoteApi) GetIndex(args ArgsId) Note {
	// Get file with product
	var item Note
	err := cmhp_file.ReadJSON(DataDir+"/note/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Note not found!")
	}
	return item
}

// Get list of all notes
func (r NoteApi) GetList() []Note {
	files, _ := cmhp_file.List(DataDir + "/note")
	out := make([]Note, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Created.Unix() > out[j].Created.Unix()
	})
	return out
}

// Create new note
func (r NoteApi) PostIndex(args Note) {
	args.Id = cmhp_crypto.UID(10)
	args.Created = time.Now()
	cmhp_file.WriteJSON(DataDir+"/note/"+args.Id+".json", &args)
}

// Update note
func (r NoteApi) PatchIndex(args Note) {
	cmhp_file.WriteJSON(DataDir+"/note/"+args.Id+".json", &args)
}

// Delete note
func (r NoteApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(DataDir + "/note/" + args.Id + ".json")
}
