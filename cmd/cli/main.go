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
	r, err := storage.FindAll(context.Background())
	if err != nil {
		fmt.Printf("failed to read %v", err)
	}
	switch command {

	case "ls":

		printState(r)
	case "shutdown":

		t := findTargetMac(r, args[1])

		if t.MacAddress == "" {
			fmt.Println("target not found")
			return
		}

		err := storage.Push(context.Background(), t.MacAddress, wakepc.NewShutdownCommand(t.MacAddress))
		if err != nil {
			fmt.Printf("failed to push message %v", err)
			return
		}
		fmt.Println("successfully sent shutdown command")

	case "wol":

		t := findTargetMac(r, args[1])

		if t.MacAddress == "" {
			fmt.Println("target not found")
			return
		}

		for _, pc := range r {

			err := storage.Push(context.Background(), pc.MacAddress, wakepc.NewShutdownCommand(t.MacAddress))
			if err != nil {
				fmt.Printf("failed to push message %v", err)
				return
			}
		}
		fmt.Println("successfully sent shutdown command")
	default:
		help()

	}
}
func help() {
	print("usage: ls|shutdown")
}
func printState(pcs []wakepc.PcState) {

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
