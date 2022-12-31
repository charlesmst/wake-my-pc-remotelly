package wakepc

import (
	"context"
	"fmt"
)

type PcStatus string

const (
	Off PcStatus = "on"
	On           = "off"
)

type PcState struct {
	MacAddress string
	State      PcStatus
}

type PcCommand string

const (
	Shutdown PcStatus = "shutdown"
	TurnOn   PcStatus = "turnon"
)

type PcStateStorage interface {
	Save(cxt context.Context, state PcState) error
	Find(ctx context.Context, mac string) (PcState, error)
	Listen(ctx context.Context, mac string, listen chan PcCommand) error
}

type PcStateStorageMock struct {
	listener chan PcCommand
	values   map[string]PcState
}

var _ PcStateStorage = &PcStateStorageMock{}

func NewPcStateStorageMock() PcStateStorageMock {
	m := make(map[string]PcState)
	return PcStateStorageMock{values: m}
}

func (p *PcStateStorageMock) Save(cxt context.Context, state PcState) error {
	p.values[state.MacAddress] = state
	return nil
}
func (p *PcStateStorageMock) Find(cxt context.Context, mac string) (PcState, error) {
	value, ok := p.values[mac]
	if !ok {
		return PcState{}, fmt.Errorf("state not found")
	}
	return value, nil
}

func (p *PcStateStorageMock) Listen(ctx context.Context, mac string, listen chan PcCommand) error {
	p.listener = listen
	return nil

}
