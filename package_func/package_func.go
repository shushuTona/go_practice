package package_func

type listSlicetype = []int

var listSlice = []int{}
func ExecuteFunc(num int) (listSlicetype, int) {
	listSlice = append(listSlice, num)

	sumResult := 0
	for _, value := range listSlice {
		sumResult += value
	}

	return listSlice, sumResult
}
