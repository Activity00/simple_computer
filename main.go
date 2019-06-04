package main

import "./bus"

func main() {
	bs := bus.NewBus()
	bs.SyncClockWire(false)

}
