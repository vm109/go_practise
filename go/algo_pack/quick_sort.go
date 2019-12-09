package algo_pack

import "fmt"

/*
quick sort works on the fact that if a number is in sorted position then that means all the
number before it are smaller than the number
and all the numbers after it are greater
[1,9,6,8,2]
in the above array 1 is in sorted position since all the numbers after it are greater than it
 */

func QuickSort(arr []int) {

	quicksort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quicksort(arr []int, l, h int){
	if l<h {
		j := partition(arr, l, h)
		fmt.Println(arr)
		fmt.Println(j)
		quicksort(arr, l, j)
			quicksort(arr, j+1, h)
	}
}
func partition(arr []int, l, h int) int{
pivot := arr[l]
i := l
j := h
for i<j{
	if arr[i] <= pivot{
		i++
	}
	if arr[j] >= pivot{
		j--
	}
	if i<j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}
	temp := arr[l]
	arr[l] = arr[j]
	arr[j] =  temp

	return j
}