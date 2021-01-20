package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Root Context
	rootCtx := context.Background()

	cancelDemo(rootCtx)
}

// START OMIT
// cancelDemo demonstrates a cancellation demo where both mainOperation
// and optionalOperation complete successfully
func cancelDemo(rootCtx context.Context) {
	cancelCtx, cancel := context.WithCancel(rootCtx)

	go optionalOperation(cancelCtx, 1 * time.Second)
	mainOperation(cancel)
}

func optionalOperation(ctx context.Context, sleepTime time.Duration) {
	select {
	case <-time.After(sleepTime):
		fmt.Println("Optional operation complete")
	case <-ctx.Done():
		fmt.Println("Early cancellation")
		return
	}
}

func mainOperation(cancel context.CancelFunc) {
	<-time.After(2 * time.Second)
	fmt.Println("Main operation complete")
	cancel()
}
// END OMIT