package wakepc

func NewShutdownCommand(mac string) PcCommandEvent {
	return PcCommandEvent{
		Command: Shutdown,
		Target:  mac,
	}
}

func NewWolCommand(mac string) PcCommandEvent {
	return PcCommandEvent{
		Command: Wol,
		Target:  mac,
	}
}
