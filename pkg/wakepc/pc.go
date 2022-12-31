package wakepc

import (
	"context"
	"log"
	"net"
	"time"
)

type Pc struct {
	Storage    PcStateStorage
	Controller PcController
	mac        string
}

func NewPcDaemon(storage PcStateStorage, controller PcController) Pc {

	mac, err := getMacAddr()
	if err != nil {
		log.Fatalf("got an error finding mac address %s", err)
	}
	return Pc{
		Storage:    storage,
		Controller: controller,
		mac:        mac[0],
	}
}

func (p *Pc) Start(ctx context.Context) {

	c := make(chan PcCommandEvent)
	err := p.Storage.Listen(ctx, p.mac, c)
	if err != nil {
		log.Fatalf("could not start listening commands")
	}
	for {
		select {
		case <-ctx.Done():
			log.Printf("stopping pc daemon")
			return
		case command := <-c:
			p.handle(ctx, command)
		case <-time.After(100 * time.Millisecond):
			p.report(ctx)
		}
	}
}
func (p *Pc) report(ctx context.Context) {
	p.Storage.Save(ctx, PcState{
		MacAddress: p.mac,
		State:      On,
	})
}

func (p *Pc) handle(ctx context.Context, command PcCommandEvent) {
	switch command.Command {
	case Shutdown:

		p.Controller.Shutdown(ctx)
	case Wol:
		p.Controller.Wol(ctx, command.Args[0])
	default:
		log.Printf("ignoring command %s", command)
	}

}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}
