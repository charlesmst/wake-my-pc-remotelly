package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/charlesmst/wake-my-pc-remotelly/pkg/backend"
	"github.com/charlesmst/wake-my-pc-remotelly/pkg/wakepc"
)

func main() {
	storage := backend.NewFirebaseStorage()

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		help()
		return
	}
	command := args[0]
	switch command {

	case "daemon":
		controller := wakepc.ResolveController()
		daemon := wakepc.NewPcDaemon(&storage, controller)
		daemon.Start(context.Background())
	case "ls":
		printState(&storage)
	case "shutdown":
		runOnTarget(&storage, args, wakepc.NewShutdownCommand)
	case "restart":
		runOnTarget(&storage, args, wakepc.NewRestartCommand)
	case "wol":
		runOnAllTargets(&storage, args, wakepc.NewWolCommand)
	default:
		help()

	}
}
func help() {
	print("usage: ls|shutdown|restart|wol [target]\n")
}
func printState(storage wakepc.PcStateStorage) {

	pcs, err := storage.FindAll(context.Background())
	if err != nil {
		fmt.Printf("failed to read %v", err)
	}

	for _, pc := range pcs {
		fmt.Printf("%s\t%s\t%s\n", pc.MacAddress, pc.HostName, pc.RealState())
	}

}

func findTargetMac(pcs []wakepc.PcState, target string) wakepc.PcState {
	for _, pc := range pcs {
		if pc.HostName == target || pc.MacAddress == target {
			return pc
		}

	}
	return wakepc.PcState{}
}

func runOnTarget(storage wakepc.PcStateStorage, args []string, commandFunc func(string) wakepc.PcCommandEvent) {

	r, err := storage.FindAll(context.Background())
	if err != nil {
		fmt.Printf("failed to read %v", err)
	}

	t := findTargetMac(r, args[1])

	if t.MacAddress == "" {
		fmt.Println("target not found")
		return
	}

	err = storage.Push(context.Background(), t.MacAddress, commandFunc(t.MacAddress))
	if err != nil {
		fmt.Printf("failed to push message %v", err)
		return
	}
	fmt.Printf("successfully sent %s command\n", args[0])

}

func runOnAllTargets(storage wakepc.PcStateStorage, args []string, commandFunc func(string) wakepc.PcCommandEvent) {

	r, err := storage.FindAll(context.Background())
	if err != nil {
		fmt.Printf("failed to read %v", err)
	}

	t := findTargetMac(r, args[1])

	if t.MacAddress == "" {
		fmt.Println("target not found")
		return
	}
	for _, pc := range r {
		err := storage.Push(context.Background(), pc.MacAddress, commandFunc(t.MacAddress))
		if err != nil {
			fmt.Printf("failed to push message %v", err)
			return
		}
	}
	fmt.Printf("successfully sent %s command to all targets\n", args[0])

}
