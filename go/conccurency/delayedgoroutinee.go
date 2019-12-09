package conccurency

import (
	"fmt"
	"math/rand"
	"time"
)

func delayedprocess(n int){
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func DelayedConcurrency(n int){
	fmt.Println("starting process")
	go  delayedprocess(n)
}