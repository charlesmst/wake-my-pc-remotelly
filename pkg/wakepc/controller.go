package wakepc

import (
	"context"
	"fmt"
	"net"
)

type PcController interface {
	FindState(ctx context.Context) (PcState, error)
	Shutdown(ctx context.Context) error
	Wol(ctx context.Context, mac string) error
}

type PcControllerMock struct {
	lastState string
	fakeMac   string
}

func (c *PcControllerMock) FindState(ctx context.Context) (PcState, error) {
	return PcState{MacAddress: c.fakeMac, State: On}, nil
}

func (c *PcControllerMock) Shutdown(ctx context.Context) error {
	c.lastState = "shutdown"
	return nil
}

func (c *PcControllerMock) Wol(ctx context.Context, mac string) error {
	c.lastState = fmt.Sprintf("wol %s", mac)
	return nil
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
