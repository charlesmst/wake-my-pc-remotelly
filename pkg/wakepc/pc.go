package wakepc

import (
	"context"
	"log"
	"time"
)

type Pc struct {
	Storage    PcStateStorage
	Controller PcController
}

func NewPcDaemon(storage PcStateStorage, controller PcController) Pc {

	return Pc{
		Storage:    storage,
		Controller: controller,
	}
}

func (p *Pc) Start(ctx context.Context) {

	initial, err := p.Controller.FindState(ctx)
	if err != nil {
		log.Fatalf("could not find initial state %v", err)
	}
	c := make(chan PcCommandEvent)

	err = p.Storage.Listen(ctx, initial.MacAddress, c)
	if err != nil {
		log.Fatalf("could not start listening commands")
	}

	p.report(ctx)
	for {
		select {
		case <-ctx.Done():
			log.Printf("stopping pc daemon")
			return
		case command := <-c:
			p.handle(ctx, command)
		case <-time.After(5000 * time.Millisecond):
			p.report(ctx)
		}
	}
}
func (p *Pc) report(ctx context.Context) {
	state, err := p.Controller.FindState(ctx)

	if err != nil {
		log.Printf("failed to get state %v", err)
		return
	}

	p.Storage.Save(ctx, state)
}

func (p *Pc) handle(ctx context.Context, command PcCommandEvent) {
	log.Printf("received command %v\n", command)
	switch command.Command {
	case Shutdown:

		p.Controller.Shutdown(ctx)
	case Wol:
		if len(command.Args) != 1 {
			return
		}
		p.Controller.Wol(ctx, command.Args[0])
	default:
		log.Printf("ignoring command %s", command)
	}

}
