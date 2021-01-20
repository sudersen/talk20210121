package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"golang.org/x/sync/errgroup"
)

const (
	userIDKey = "user_id"
)

func main() {
	// Root Context
	rootCtx := context.Background()
	rand.Seed(time.Now().UnixNano())

	ctx := NewContext(rootCtx, fmt.Sprintf("abcd%d", rand.Int() % 10000))
	fmt.Println("set user-id is", ctx.Value(userIDKey))
}

// START OMIT
// NewContext demonstrates a way of using WithValue
func NewContext(rootCtx context.Context, userID string) context.Context {
	return context.WithValue(rootCtx, userIDKey, userID)
}
// END OMIT