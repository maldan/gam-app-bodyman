package caloryman

import (
	"strings"

	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"github.com/maldan/go-restserver"
)

type ProductApi struct {
}

// Get product without error
func (r ProductApi) GetSafeIndex(args ArgsId) Product {
	// Get file with product
	var item Product
	cmhp_file.ReadJSON(DataDir+"/product/"+args.Id+".json", &item)
	return item
}

// Get product
func (r ProductApi) GetIndex(args ArgsId) Product {
	// Get file with product
	var item Product
	err := cmhp_file.ReadJSON(DataDir+"/product/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	return item
}

// Get product list
func (r ProductApi) GetList() []Product {
	files, _ := cmhp_file.List(DataDir + "/product")
	out := make([]Product, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	return out
}

// Get product by name
func (r ProductApi) GetByName(args ArgsName) Product {
	var productList = r.GetList()
	product, itemId := cmhp_slice.FindR(productList, func(i interface{}) bool {
		return strings.Contains(strings.ToLower(i.(Product).Name), strings.ToLower(args.Name))
	})
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	return product.(Product)
}

// Add new product
func (r ProductApi) PostIndex(args Product) {
	args.Id = cmhp_crypto.UID(10)
	args.Protein = UCTo(UCFrom(args.Protein), "g")
	args.Fat = UCTo(UCFrom(args.Fat), "g")
	args.Carbohydrate = UCTo(UCFrom(args.Carbohydrate), "g")
	args.ComponentIdList = make([]string, 0)
	cmhp_file.WriteJSON(DataDir+"/product/"+args.Id+".json", &args)
}

// Update product
func (r ProductApi) PatchIndex(args Product) {
	args.Protein = UCTo(UCFrom(args.Protein), "g")
	args.Fat = UCTo(UCFrom(args.Fat), "g")
	args.Carbohydrate = UCTo(UCFrom(args.Carbohydrate), "g")
	cmhp_file.WriteJSON(DataDir+"/product/"+args.Id+".json", &args)
}

// Delete product
func (r ProductApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(DataDir + "/product/" + args.Id + ".json")
}
