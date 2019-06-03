package bus

type MasterMux struct {
}

func NewMasterMux(bus *Bus) *MasterMux {
	return &MasterMux{}
}
