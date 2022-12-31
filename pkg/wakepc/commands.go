package wakepc

func NewShutdownCommand(mac string) PcCommandEvent {
	return PcCommandEvent{
		Command: Shutdown,
		Target:  mac,
	}
}
