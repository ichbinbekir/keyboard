package keyboard

type Event interface{}

type KeyboardEvent struct {
	Key   uint32
	State bool
}

// TODO:
// type MouseEvent struct {
// 	Event
// }
