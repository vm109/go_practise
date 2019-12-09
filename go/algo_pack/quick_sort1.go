package algo_pack

import "fmt"

func QuickSort1() {
	fmt.Println(partition1([]int{11, 2, 3, 1, 6,21,0,41,5}, 0))
}

func partition1(arr1 []int, j int) []int {
	if j < len(arr1){
	arr := arr1[0:len(arr1)-j]
	pivot := arr[len(arr)-1]
		var l, r []int
		for i := 0; i < len(arr); i++ {
			if arr[i] < pivot {
				l = append(l, arr[i])
			} else if arr[i] > pivot {
				r = append(r, arr[i])
			}
		}
		arr = append(append(l, pivot), r...)
		arr1 = append(arr,arr1[len(arr1)-j: len(arr1)]...)

		arr1 = partition1(arr1, j+1)

	}
	return arr1
}
