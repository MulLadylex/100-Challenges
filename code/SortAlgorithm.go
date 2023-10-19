package code

/*
 冒泡排序，对给定的数组进行的
*/
func bubbleSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1] = arr[j] + arr[j+1]
				arr[j] = arr[j+1] - arr[j]
				arr[j+1] = arr[j+1] - arr[j]
			}
		}
	}
	for _, v := range arr {
		print(v, " ")
	}
}

func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}
