package caloryman

import (
	"strings"

	"github.com/maldan/go-cmhp"
	"github.com/maldan/go-docdb"
	"github.com/maldan/go-restserver"
	"github.com/rs/xid"
)

type ProductApi int

type ProductApi_PostIndexArgs struct {
	Id           string
	Name         string
	Fat          string
	Protein      string
	Carbohydrate string
}

type DeleteIndexArgs struct {
	Id string
}

func (f ProductApi) GetIndex(args IdArgs) Product {
	var product []Product
	docdb.Get("db", "product", &product)
	item, itemId := cmhp.SliceFindR(product, func(i interface{}) bool {
		return i.(Product).Id == args.Id
	})
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	return item.(Product)
}

func (f ProductApi) GetList() []Product {
	var product []Product
	docdb.Get("db", "product", &product)
	return product
}

func (f ProductApi) GetByName(args NameArgs) Product {
	var productList = f.GetList()
	product, itemId := cmhp.SliceFindR(productList, func(i interface{}) bool {
		return strings.Contains(strings.ToLower(i.(Product).Name), strings.ToLower(args.Name))
	})
	if itemId == -1 {
		restserver.Error(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	return product.(Product)
}

func (f ProductApi) PostIndex(args ProductApi_PostIndexArgs) {
	var productList = f.GetList()
	productList = append(productList, Product{
		Id:              xid.New().String(),
		Name:            args.Name,
		Protein:         args.Protein,
		Fat:             args.Fat,
		Carbohydrate:    args.Carbohydrate,
		ComponentIdList: make([]string, 0),
	})
	docdb.Save("db", "product", &productList)
}

func (f ProductApi) PatchIndex(args ProductApi_PostIndexArgs) {
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
}