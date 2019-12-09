package algo_pack

import "fmt"

/*
Merge sort:
recursively check if two numbers and by comparison swap those numbers
*/

func MergerSort(arr []int) {
	var length int = len(arr)
	midpoint := length / 2
	larr := arr[:midpoint]
	rarr := arr[midpoint:length]
	recursiveCompare(larr, rarr, []int{})
}

func recursiveCompare(larr, rarr []int, actual_arr []int) {
	if len(larr) > 1 && len(rarr) > 1 {
		var left_arr_lem int = len(larr)
		left_arr_left := larr[:left_arr_lem/2]
		left_arr_right := larr[left_arr_lem/2 : left_arr_lem]
		var right_arr_lem int = len(rarr)
		right_arr_left := larr[:right_arr_lem/2]
		right_arr_right := larr[right_arr_lem/2 : right_arr_lem]
		recursiveCompare(left_arr_left, left_arr_right, actual_arr)
		recursiveCompare(right_arr_left, right_arr_right, actual_arr)
	}
if len(larr)<=1 && len(rarr)<=1 {
 fmt.Println(larr)
 fmt.Println(rarr)
}

}
