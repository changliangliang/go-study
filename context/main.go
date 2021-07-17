package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context, num int) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("任务[%d]运行结束\n", num)
			return
		default:
			fmt.Println(ctx.Value("key"))
			fmt.Printf("任务[%d]执行中......\n", num)
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	ctx = context.WithValue(ctx, "key", "value")
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	for i := 0; i < 5; i++ {
		go task(ctx, i)
	}
	time.Sleep(10 * time.Second)

}
