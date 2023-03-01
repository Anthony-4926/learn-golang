package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second))
	ctx.Done()

	fmt.Println(ctx.Deadline())
}
