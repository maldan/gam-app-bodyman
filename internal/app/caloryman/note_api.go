package caloryman

import (
	"sort"
	"time"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type NoteApi int

type NoteApi_PostIndexArgs struct {
	Id          string
	Description string
	Created     time.Time
}

// Get note by id
func (f NoteApi) GetIndex(args IdArgs) (Note, int) {
	noteList := f.GetList()
	item, itemId := cmhp.SliceFindR(noteList, func(i interface{}) bool {
		return i.(Note).Id == args.Id
	})

	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Note not found!")
	}
	return item.(Note), itemId
}

// Get list of all notes
func (f NoteApi) GetList() []Note {
	var noteList []Note
	docdb.Get("db", "note", &noteList)
	sort.SliceStable(noteList, func(i, j int) bool {
		return noteList[i].Created.Unix() > noteList[j].Created.Unix()
	})
	return noteList
}

// Create new note
func (f NoteApi) PostIndex(args NoteApi_PostIndexArgs) {
	var noteList = f.GetList()
	noteList = append(noteList, Note{
		Id:          xid.New().String(),
		Description: args.Description,
		Created:     time.Now(),
	})
	docdb.Save("db", "note", &noteList)
}

// Update note
func (f NoteApi) PatchIndex(args NoteApi_PostIndexArgs) {
	note, noteIndex := f.GetIndex(IdArgs{Id: args.Id})
	noteList := f.GetList()
	note.Description = args.Description
	noteList[noteIndex] = note
	docdb.Save("db", "note", &noteList)
}

func (f NoteApi) DeleteIndex(args DeleteIndexArgs) {
	noteList := f.GetList()
	out := cmhp.SliceFilterR(noteList, func(i interface{}) bool {
		return i.(Note).Id != args.Id
	})
	docdb.Save("db", "note", &out)
}