package main

import(
	"context"
	"time"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	sem := make(chan struct{},3)
	for i:=0; i<10; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			sem <- struct{}{}
			defer func(){<-sem} ()
			worker(ctx, i)
		}(i)
	}
	wg.Wait()
}

func worker(ctx context.Context, id int){
	select{
	case <-ctx.Done():
		fmt.Println("worker %d cancled", id)
		return
	default:
		fmt.Println("worker %d working", id)
		time.Sleep(1*time.Second)
	}
}
  