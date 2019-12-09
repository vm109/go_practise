package slice_pack

func CreateNewSlice(len, capacity int)[]int{
	return make([]int, len, capacity)
}
