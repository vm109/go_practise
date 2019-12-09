package algo_pack

/*
Bubble sort - it is the simple sorting algorithm,
	compares two adjacent numbers and then swaps them
	in either ascending or descending order.

Continue this process until the entire array is sorted
 */


func BubbleSort(arr []int) []int {
	//for i,num := range arr{
	//	fmt.Println("index is "+ strconv.Itoa(i)+ " number is "+ strconv.Itoa(num) )
	//	fmt.Println(arr[i+1])
	//	fmt.Println("----------")
	//}
	var swapflag int8 = 1
for swapflag >0 {
	swapflag = 0
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = temp
			swapflag = swapflag+1
		}
	}
}
return arr
}