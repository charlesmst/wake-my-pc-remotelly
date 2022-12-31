package main

import (
	"context"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := wakepc.NewPcStateStorageMock()
	controller := wakepc.PcControllerMock{}
	daemon := wakepc.NewPcDaemon(&storage, &controller)
	daemon.Start(context.Background())

}
