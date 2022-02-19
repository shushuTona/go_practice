package package_var

import "fmt"

type itemType = int

func PrintListSlice() {
	var listSlice = []itemType{}

	listSlice = append(listSlice, 100);
	listSlice = append(listSlice, 200);
	listSlice = append(listSlice, 300);
	
	fmt.Println(listSlice)
	fmt.Println(len(listSlice))
	fmt.Println(cap(listSlice))
}

func PrintMap() {
	var mapList = map[string]int{
		"hoge": 100,
		"piyo": 200,
		"fuga": 300,
	}
	fmt.Println(mapList)

	mapList["hoge"] = 999
	fmt.Println(mapList)

	delete(mapList, "fuga")
	fmt.Println(mapList)

	fmt.Println(len(mapList))

	// mapに保存されている値、mapに存在するかの真偽値を取得する
	_, is_exist := mapList["hoge"]
	if is_exist {
		fmt.Println("exist hoge")
	} else {
		fmt.Println("not exist hoge")
	}

	for key, value := range mapList {
		fmt.Println(key, value)
	}
}
