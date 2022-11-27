package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}
func main() {
	ctx := context.Background()
	process(ctx)

	ctx = context.WithValue(ctx, "traceId", "xjl")
	process(ctx)
}
