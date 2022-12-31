package main

import (
	"context"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := wakepc.PcStateStorageMock{}
	daemon := wakepc.NewPcDaemon(&storage)
	daemon.Start(context.Background())

}
