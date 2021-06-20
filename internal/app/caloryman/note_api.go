package caloryman

import (
	"time"

	"github.com/maldan/go-docdb"
	"github.com/rs/xid"
)

type NoteApi int

type NoteApi_PostIndexArgs struct {
	Id          string
	Description string
	Created     time.Time
}

func (f NoteApi) GetList() []Note {
	var noteList []Note
	docdb.Get("db", "note", &noteList)
	return noteList
}

func (f NoteApi) PostIndex(args NoteApi_PostIndexArgs) {
	var noteList = f.GetList()
	noteList = append(noteList, Note{
		Id:      xid.New().String(),
		Created: time.Now(),
	})
	docdb.Save("db", "note", &noteList)
}

/*func (f ProductApi) PatchIndex(args ProductApi_PostIndexArgs) {
	var product []Product
	docdb.Get("db", "product", &product)
	item, itemId := cmhp.SliceFindR(product, func(i interface{}) bool {
		return i.(Product).Id == args.Id
	})

	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	productItem := item.(Product)
	productItem.Name = args.Name
	product[itemId] = productItem

	docdb.Save("db", "product", &product)
}

func (f ProductApi) DeleteIndex(args DeleteIndexArgs) {
	var productList = f.GetList()
	out := cmhp.SliceFilterR(productList, func(i interface{}) bool {
		return i.(Product).Id != args.Id
	})
	docdb.Save("db", "product", &out)
}*/
