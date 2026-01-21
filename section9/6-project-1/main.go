// package main

// import (
// 	"fmt"
// 	"time"
// )

// func ping(ch chan string) {
// 	for {
// 		msg := <-ch // received msg from channel
// 		fmt.Println("ping received:", msg)
// 		time.Sleep(time.Second)
// 		ch <- "pong" // send msg to channel
// 	}
// }

// func pong(ch chan string) {
// 	for {
// 		msg := <-ch // received msg from channel
// 		fmt.Println("pong received:", msg)
// 		time.Sleep(time.Second)
// 		ch <- "ping" // send msg to channel
// 	}
// }

// func main() {
// 	ch := make(chan string)

// 	go ping(ch)
// 	go pong(ch)

// 	ch <- "ping" // send msg ping for the first time

// 	time.Sleep(5 * time.Second)
// }

package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string, done chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			close(done)
			return
		case msg := <-ch:
			fmt.Println("ping received:", msg)
			time.Sleep(1 * time.Second)
			ch <- "pong"
		}
	}
}

func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ch:
			fmt.Println("pong received:", msg)
			time.Sleep(1 * time.Second)
			ch <- "ping"
		}
	}
}

func main() {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	ch := make(chan string)
	done := make(chan struct{})

	go ping(ctx, ch, done)
	go pong(ctx, ch)

	ch <- "ping"

	<-done
	fmt.Println("finished")

}
