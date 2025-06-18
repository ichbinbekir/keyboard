package user32

import (
	"syscall"
)

var _user32 = syscall.NewLazyDLL("user32.dll")

var (
	_setWindowsHookExW   = _user32.NewProc("SetWindowsHookExW")
	_getMessageW         = _user32.NewProc("GetMessageW")
	_translateMessage    = _user32.NewProc("TranslateMessage")
	_dispatchMessageW    = _user32.NewProc("DispatchMessageW")
	_unhookWindowsHookEx = _user32.NewProc("UnhookWindowsHookEx")
	_callNextHookEx      = _user32.NewProc("CallNextHookEx")
	_getAsyncKeyState    = _user32.NewProc("GetAsyncKeyState")
	_keybd_event         = _user32.NewProc("keybd_event")
)
