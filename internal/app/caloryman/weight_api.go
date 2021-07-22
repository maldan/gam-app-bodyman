package caloryman

type WeightApi struct {
}

/*func (f WeightApi) GetIndex(args ArgsId) (Weight, int) {
	// Find training
	list := f.GetList()
	item, itemId := cmhp.SliceFindR(list, func(i interface{}) bool {
		return i.(Weight).Id == args.Id
	})

	// Not found
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Weight not found!")
	}

	return item.(Weight), itemId
}

func (f WeightApi) GetList() []Weight {
	var list []Weight
	docdb.Get(DataDir, f.Table, &list)
	return list
}

func (f WeightApi) PostIndex(args Weight) {
	list := f.GetList()
	args.Id = xid.New().String()
	list = append(list, args)
	docdb.Save(DataDir, f.Table, &list)
}

func (f WeightApi) PatchIndex(args Weight) {
	_, itemId := f.GetIndex(ArgsId{Id: args.Id})
	list := f.GetList()
	list[itemId] = args
	docdb.Save(DataDir, f.Table, &list)
}

func (f WeightApi) DeleteIndex(args ArgsId) {
	list := f.GetList()
	out := cmhp.SliceFilterR(list, func(i interface{}) bool {
		return i.(Weight).Id != args.Id
	})
	docdb.Save(DataDir, f.Table, &out)
}*/
