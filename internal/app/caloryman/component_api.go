package caloryman

type ComponentApi int

type ComponentApi_PostIndexArgs struct {
	Id          string
	Name        string
	Type        string
	Description string
}

/*func (f ComponentApi) GetIndex(args ArgsId) Component {
	var component []Component
	docdb.Get(DataDir, "component", &component)
	item, itemId := cmhp.SliceFindR(component, func(i interface{}) bool {
		return i.(Component).Id == args.Id
	})
	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Component not found!")
	}
	return item.(Component)
}

func (f ComponentApi) GetList() []Component {
	var component []Component
	docdb.Get(DataDir, "component", &component)
	return component
}

func (f ComponentApi) PostIndex(args ComponentApi_PostIndexArgs) {
	var component []Component
	docdb.Get(DataDir, "component", &component)
	component = append(component, Component{
		Id:          xid.New().String(),
		Name:        args.Name,
		Type:        args.Type,
		Description: args.Description,
	})
	docdb.Save(DataDir, "component", &component)
}

func (f ComponentApi) PatchIndex(args ComponentApi_PostIndexArgs) {
	var component []Component
	docdb.Get(DataDir, "component", &component)
	item, itemId := cmhp.SliceFindR(component, func(i interface{}) bool {
		return i.(Component).Id == args.Id
	})

	if itemId == -1 {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Component not found!")
	}
	componentItem := item.(Component)
	componentItem.Name = args.Name
	componentItem.Type = args.Type
	componentItem.Description = args.Description
	component[itemId] = componentItem

	docdb.Save(DataDir, "component", &component)
}

func (f ComponentApi) DeleteIndex(args ArgsId) {
	var component []Component
	docdb.Get(DataDir, "component", &component)
	out := cmhp.SliceFilterR(component, func(i interface{}) bool {
		return i.(Component).Id != args.Id
	})
	docdb.Save(DataDir, "component", &out)
}*/
