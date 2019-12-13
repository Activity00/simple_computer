package io

type DisplayAdapter struct {
}

func NewDisplaydAdapter() *DisplayAdapter {
	d := new(DisplayAdapter)
	return d
}

type ScreenControl struct {
	adapter *DisplayAdapter
}

func NewScreenControl(adapter *DisplayAdapter, outputChan chan *[160][240]byte, quit chan bool) *ScreenControl {
	s := new(ScreenControl)
	s.adapter = adapter
	return s
}
