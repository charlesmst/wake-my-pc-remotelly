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

	case "ls":

		r, err := storage.FindAll(context.Background())
		if err != nil {
			fmt.Printf("failed to read %v", err)
		}
		printState(r)
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
