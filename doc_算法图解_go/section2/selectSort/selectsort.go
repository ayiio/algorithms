package selectsort

func findSmaller(arr []int) (smaller_index int) {
	smalller := arr[0]  //最小值
	smaller_index = 0   //最小值索引
	for i := 1; i < len(arr); i++ {
		if arr[i] < smalller {
			smalller = arr[i]
			smaller_index = i
		}
	}
	return
}

func SelectSort(arr []int) []int {
	var newArr []int
	leno := len(arr)
	for i := 0; i < leno; i++ {
		smallidx := findSmaller(arr)  //找到数组中最小元素，将其加入新数组中
		newArr = append(newArr, arr[smallidx])
		arr = append(arr[:smallidx], arr[smallidx+1:]...)
	}
	return newArr
}
