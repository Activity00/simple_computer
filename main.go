package main

import "simple_computer/computer"

func main() {
	screenChannel := make(chan *[160][240]byte)
	quitChannel := make(chan bool, 10)
	simpleComputer := computer.NewComputer(screenChannel, quitChannel)
	simpleComputer.PowerUp()
}
