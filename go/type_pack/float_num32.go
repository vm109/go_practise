package type_pack

import "math"

func GetFloat_Number_32() float32 {
	var number = float32(math.Sin(0.5))
	return number
}


func Float32_Sum_Two_Num(num1 , num2 float32) float32{
	return num1+num2
}