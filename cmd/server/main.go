package main

import (
	"context"
	"fmt"
	"trinity-be/global"
	"trinity-be/internal/initialize"
)

func main() {
    initialize.InitializeServer()
	ctx := context.Background()
    global.RedisDB.Set(ctx, "is_running", "OK", 600)
    fmt.Println(global.RedisDB.Get(ctx, "is_running_2"))
}
