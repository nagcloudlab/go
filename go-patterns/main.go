package main

import "time"

/*
	main-component / subject  ( e.g Door )
		- open()
		- close()
	listener / observer
		- light ( e.g. light on/off)
		- AC ( e.g. turn on/off)
		- Fan ( e.g. turn on/off)
*/

// Observer pattern

type DoorListener interface {
	On()
	Off()
}

type Light struct{}

func (l *Light) On() {
	println("Light is on")
}
func (l *Light) Off() {
	println("Light is off")
}

type Door struct {
	listeners []DoorListener
}

func (d *Door) AddListener(listener DoorListener) {
	d.listeners = append(d.listeners, listener)
}
func (d *Door) RemoveListener(listener DoorListener) {
	for i, l := range d.listeners {
		if l == listener {
			d.listeners = append(d.listeners[:i], d.listeners[i+1:]...)
			break
		}
	}
}

func (d *Door) Open() {
	println("Door is open")
	for _, listener := range d.listeners {
		listener.On()
	}
}
func (d *Door) Close() {
	println("Door is closed")
	for _, listener := range d.listeners {
		listener.Off()
	}
}

func main() {
	door := &Door{}
	light := &Light{}
	door.AddListener(light)

	door.Open()
	time.Sleep(1 * time.Second)
	door.Close()

}
