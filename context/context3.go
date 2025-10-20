package main

import(
	"context"
	"time"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "worker1")
	go worker(ctx, "worker2")

	time.Sleep(1*time.Second)
	cancel()
	time.Sleep(1*time.Second)
}

func worker(ctx context.Context, name string){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("stoped",name)
			return
		default :
			fmt.Println("working", name)
			time.Sleep(500*time.Millisecond)
		}
	}
}