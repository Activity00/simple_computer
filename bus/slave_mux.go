package bus

type SlaveMux struct {
}

func NewSlaveMux(bus *Bus) *SlaveMux {
	return &SlaveMux{}
}
