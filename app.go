package main

import "./go/algo_pack"
func main() {
 algo_pack.QuickSort1()
}




//math_p.Math_calculation()
//math_p.Fun()
//fmt.Printf("the number is %v",type_pack.GetFloat_Number_32())
//fmt.Printf("the number is %v",type_pack.GetIntNumber32())

// summing two float32 numbers
//var x, y float32 = 1.002, 2.004
//
//fmt.Println(type_pack.Float32_Sum_Two_Num(x, y))
//// it throws an error because go goes for default type of float64
//p, q := 2.009, 3.006
//fmt.Println(type_pack.Float32_Sum_Two_Num(p, q))

// String
//mssg1, mssg2 := "dooms day dooms day","we are on war"
//var  enhanced1, enhanced2 string= type_pack.Get_Enhanced_String(mssg1, mssg2)
// fmt.Println(enhanced1)
// fmt.Println(enhanced2)

//Pointers
//number will have the number 15 as value
//number := 15
//// pointer will have the address of the number as value
//pointer := &number
//
//fmt.Printf("the address of the number is %v \n", pointer)
//fmt.Printf("the value is %v \n",*pointer)
//fmt.Printf("I am wondering what would be the value of %v \n",*&number)
//
////now lets change the value at the adderss
//*pointer = 7
//
//fmt.Printf("this will change the number value from 15 to %v \n", number)
//// so basically the above statement makes it clear that *points is equal to number
//
//// what would this print?
//fmt.Printf("%v \n", *pointer**pointer)

//var number  = 2
//pointerpck.Add_Two(&number)
//
//fmt.Println(number)

//http.HandleFunc("/", servicepck.Greet_Server_Home)
//http.ListenAndServe(":8000", nil)
//
//teslaModex := customStructs.Create_Car(8,6,false, true)
//fmt.Println(teslaModex.CarAdvertise("tesla modelX"))
//teslaModex.ChangeCarWheelsOrSeats(4,5)
//fmt.Println(teslaModex.CarAdvertise("new tesla modelX"))

// capacity is the size of the array after the length of array
// i.e the array is now [0,0] but if the capacity is 3, the array can be extended to length of 3
//slice := slice_pack.CreateNewSlice(2,3)
//fmt.Println(slice)
//slice = append(slice, 0)
//fmt.Println(slice)
//slice = append(slice, 0)
//fmt.Println(slice)

//var s1 i.Shape = i.Rectangle{ Length:20, Width:30}
//fmt.Println(s1.Area())
//fmt.Println(s1.Perimeter())


//var building interfaces.Construction = &interfaces.Building{
//	Base:  20,
//	Paint: "berger",
//}
//
//fmt.Println(building.BuildBase())
//fmt.Println(building.PaintHouse())
//fmt.Println(building.GetConstructionCost())

//conccurency.Runconccurently(6)
//conccurency.DelayedConcurrency(8)