package main

type (
	SimpleEvent func()
)

var actions struct {
	stopping bool
	enter    SimpleEvent
	up       SimpleEvent
	down     SimpleEvent
}
