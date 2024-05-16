package main

import (
	"context"
	"fmt"
	"lore-keeper-be/internal"
)

func main() {
	//TODO use cancelcontext for something?
	fmt.Println("Starting app")
	ctx, _ := context.WithCancel(context.Background())
	app := internal.NewApp(ctx)

	app.Start()

	//TODO listen for shutdown signal
}
