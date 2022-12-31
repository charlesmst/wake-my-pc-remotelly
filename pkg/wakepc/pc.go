package wakepc

import (
	"context"
	"log"
	"net"
)

type Pc struct {
	Storage PcStateStorage
	mac     string
}

func NewPcDaemon(storage PcStateStorage) Pc {

	mac, err := getMacAddr()
	if err != nil {
		log.Fatalf("got an error finding mac address %s", err)
	}
	return Pc{
		Storage: storage,
		mac:     mac[0],
	}
}
func (p *Pc) Start(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopping pc daemon")
			return
		default:
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
