package main

import (
	"context"
	"log"
	"runtime"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/backend"
	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := backend.NewFirebaseStorage()

	controller := resolveController()

	daemon := wakepc.NewPcDaemon(&storage, controller)
	daemon.Start(context.Background())

}

func resolveController() wakepc.PcController {
	switch runtime.GOOS {
	case "linux":
		return &wakepc.LinuxController{}

	case "darwin":
		return &wakepc.MacController{}
	default:
		log.Fatalf("unsupported OS")
		return nil
	}
}
