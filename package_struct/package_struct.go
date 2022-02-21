package package_struct

type listItemType = map[string]string
type registerItemListType = map[int]listItemType

type Register struct {
	RegisterId int
	RegisterItemList registerItemListType
}

func ( register *Register ) CountUpId() {
	register.RegisterId++
}

func ( register *Register ) SetItem(listItem listItemType) registerItemListType {
	index := register.RegisterId

	register.RegisterItemList[index] = listItem

	register.CountUpId()

	return register.RegisterItemList
}
