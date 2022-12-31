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

type PcStateStorage interface {
	Save(cxt context.Context, state PcState) error
	Find(ctx context.Context, mac string) (PcState, error)
}

type PcStateStorageMock struct {
	values map[string]PcState
}

var _ PcStateStorage = &PcStateStorageMock{}

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
