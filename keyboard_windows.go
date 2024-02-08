package keyboard

import (
	"syscall"
)

var _user32 = syscall.NewLazyDLL("user32.dll")

var (
	_getAsyncKeyState = _user32.NewProc("GetAsyncKeyState")
	_mouse_event      = _user32.NewProc("mouse_event")
	_keybd_event      = _user32.NewProc("keybd_event")
	//_sendInput        = _user32.NewProc("SendInput")
)

const errOperationComletedText = "The operation completed successfully."

func getKeyState(code uintptr) (bool, error) {
	state, _, err := _getAsyncKeyState.Call(code)
	if err.Error() != errOperationComletedText {
		return false, err
	}

	return state != 0, nil
}
