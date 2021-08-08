package caloryman

var productListCache = make([]Product, 0)

/*func DB_eatList() []Eat {
	var list []Eat
	cmhp.FileReadAsJSON(DataDir+"/eat.json", &list)
	return list
}

func DB_productList() []Product {
	if len(productListCache) > 0 {
		return productListCache
	}

	var list []Product
	cmhp.FileReadAsJSON(DataDir+"/product.json", &list)
	productListCache = list
	return list
}

func DB_productById(id string) Product {
	var list = DB_productList()
	for _, item := range list {
		if item.Id == id {
			return item
		}
	}
	return Product{}
}*/
