package algo_pack

import "fmt"

func MergeTwoArray(arr1, arr2 []int){
	fmt.Println(arr1)
	fmt.Println(arr2)
	var mergeArray []int
	j := 0
	i := 0
	for !(i>= len(arr1) || j >= len(arr2)){
		arr1Element := arr1[i]
		arr2Element := arr2[j]
		if arr1Element < arr2Element {
			mergeArray = append(mergeArray, arr1Element)
			i++
		}else {
			mergeArray = append(mergeArray, arr2Element)
			j++
		}
	}
	fmt.Println("reaching here", i, j, len(arr1), len(arr2))
	if i< len(arr1) {
		for i< len(arr1) {
			mergeArray = append(mergeArray, arr1[i])
			i++
		}
	}
	if j< len(arr2) {
		for j < len(arr2) {
			mergeArray = append(mergeArray, arr2[j])
			j++
		}
	}

	fmt.Println(mergeArray)

}
