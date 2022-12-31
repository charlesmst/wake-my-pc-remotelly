package main

import (
	"context"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := wakepc.NewPcStateStorageMock()
	daemon := wakepc.NewPcDaemon(&storage)
	daemon.Start(context.Background())

}
