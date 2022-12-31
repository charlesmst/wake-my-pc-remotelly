package main

import (
	"context"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/backend"
	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	// storage := wakepc.NewPcStateStorageMock()
	// controller := wakepc.PcControllerMock{}
	// daemon := wakepc.NewPcDaemon(&storage, &controller)
	// daemon.Start(context.Background())
	// c := wakepc.LinuxController{}
	// err := c.Shutdown(context.Background())
	// log.Printf("%v", err)
	storage := backend.NewFirebaseStorage()

	controller := wakepc.LinuxController{}

	daemon := wakepc.NewPcDaemon(&storage, &controller)
	daemon.Start(context.Background())

}
