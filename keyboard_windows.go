package keyboard

import (
	"syscall"
)

var _user32 = syscall.NewLazyDLL("user32.dll")

var _getAsyncKeyState = _user32.NewProc("GetAsyncKeyState")

var errOperationComletedText = "The operation completed successfully."

func getKeyState(code uintptr) (bool, error) {
	state, _, err := _getAsyncKeyState.Call(code)
	if err.Error() != "The operation completed successfully." {
		return false, err
	}

	return state != 0, nil
}
