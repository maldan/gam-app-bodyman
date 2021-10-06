package api

import (
	"strings"

	"github.com/maldan/gam-app-caloryman/internal/app/caloryman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"github.com/maldan/go-restserver"
)

type ProductApi struct {
}

// Get product without error
func (r ProductApi) GetSafeIndex(args ArgsId) core.Product {
	// Get file with product
	var item core.Product
	cmhp_file.ReadJSON(core.DataDir+"/product/"+args.Id+".json", &item)
	return item
}

// Get product
func (r ProductApi) GetIndex(args ArgsId) core.Product {
	// Get file with product
	var item core.Product
	err := cmhp_file.ReadJSON(core.DataDir+"/product/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	return item
}

// Get product list
func (r ProductApi) GetList() []core.Product {
	files, _ := cmhp_file.List(core.DataDir + "/product")
	out := make([]core.Product, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	return out
}

// Get product by name
func (r ProductApi) GetByName(args ArgsName) core.Product {
	var productList = r.GetList()
	product, itemId := cmhp_slice.FindR(productList, func(i interface{}) bool {
		return strings.Contains(strings.ToLower(i.(core.Product).Name), strings.ToLower(args.Name))
	})
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Product not found!")
	}
	return product.(core.Product)
}

// Add new product
func (r ProductApi) PostIndex(args core.Product) {
	args.Id = cmhp_crypto.UID(10)
	args.Protein = core.UCTo(core.UCFrom(args.Protein), "g")
	args.Fat = core.UCTo(core.UCFrom(args.Fat), "g")
	args.Carbohydrate = core.UCTo(core.UCFrom(args.Carbohydrate), "g")
	args.ComponentIdList = make([]string, 0)
	cmhp_file.WriteJSON(core.DataDir+"/product/"+args.Id+".json", &args)
}

// Update product
func (r ProductApi) PatchIndex(args core.Product) {
	args.Protein = core.UCTo(core.UCFrom(args.Protein), "g")
	args.Fat = core.UCTo(core.UCFrom(args.Fat), "g")
	args.Carbohydrate = core.UCTo(core.UCFrom(args.Carbohydrate), "g")
	cmhp_file.WriteJSON(core.DataDir+"/product/"+args.Id+".json", &args)
}

// Delete product
func (r ProductApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/product/" + args.Id + ".json")
}
