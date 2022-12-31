package main

import (
	"context"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/backend"
	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := backend.NewFirebaseStorage()

	controller := wakepc.ResolveController()

	daemon := wakepc.NewPcDaemon(&storage, controller)
	daemon.Start(context.Background())

}
