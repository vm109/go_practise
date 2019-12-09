package interfaces

import "strconv"

type Construction interface {
	BuildBase() string
	PaintHouse() string
	GetConstructionCost() int32
}


type Building struct{
	Base int32
	Paint string
}

func (b *Building) BuildBase() string{
	return "construct building with base at feet  "+ strconv.Itoa(int(b.Base))
}

func (b *Building) PaintHouse() string{
	return "Painting building with the following paint "+ b.Paint
}

func (b *Building) GetConstructionCost() int32{
	return  b.Base*(int32(len(b.Paint)))*10
}