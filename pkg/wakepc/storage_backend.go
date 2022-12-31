package wakepc

import (
	"context"
	"fmt"
	"time"
)

type PcStatus string

const (
	Off PcStatus = "off"
	On           = "on"
)

type PcState struct {
	MacAddress string
	HostName   string
	State      PcStatus
	ReadTime   int64
}

func (p *PcState) RealState() PcStatus {
	if p.State != On {
		return p.State
	}

	considerOn := time.Now().Add(time.Minute * -1).Unix()
	if p.ReadTime < considerOn {
		return Off
	}
	return p.State
}

type PcCommand string

const (
	Shutdown PcCommand = "shutdown"
	Wol      PcCommand = "wol"
	Restart  PcCommand = "restart"
)

type PcCommandEvent struct {
	Command PcCommand
	Args    []string
	Target  string
	Time    int64
}

type PcStateStorage interface {
	Save(cxt context.Context, state PcState) error
	Find(ctx context.Context, mac string) (PcState, error)
	FindAll(ctx context.Context) ([]PcState, error)
	Listen(ctx context.Context, mac string, listen chan PcCommandEvent) error
	Push(ctx context.Context, mac string, event PcCommandEvent) error
}

type PcStateStorageMock struct {
	listener chan PcCommandEvent
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

func (p *PcStateStorageMock) FindAll(ctx context.Context) ([]PcState, error) {
	l := make([]PcState, 0)
	for _, value := range p.values {
		l = append(l, value)
	}
	return l, nil
}

func (p *PcStateStorageMock) Listen(ctx context.Context, mac string, listen chan PcCommandEvent) error {
	p.listener = listen
	return nil
}

func (p *PcStateStorageMock) Push(ctx context.Context, mac string, event PcCommandEvent) error {
	return nil
}
