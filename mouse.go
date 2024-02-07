package keyboard

func MouseEvent(event uintptr) {
	_mouse_event.Call(event)
}
