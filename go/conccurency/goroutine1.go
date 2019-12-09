package conccurency
/*
Concurrency - multiple process will share a work. So a task is finished only after all the
				threads complete.
Parallelism - runs multiple process at a time.
 */
import (
	"fmt"
	"strconv"
	"sync"
)

func process(loopnumber int, waitgroup *sync.WaitGroup){
	for i:=0; i<loopnumber; i++{
		fmt.Println("---"+strconv.Itoa(i))
	}
	waitgroup.Done()
}

func Runconccurently(n int){
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go process(n, &waitgroup)
	waitgroup.Wait()
	fmt.Println("process started")
}