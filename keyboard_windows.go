package keyboard

import (
	"unsafe"

	"github.com/ichbinbekir/windows"
	"github.com/ichbinbekir/windows/user32"
)

func (kb *Keyboard) readEvents() {
	hook := user32.SetWindowsHookExW(
		user32.WH_KEYBOARD_LL,
		func(code int, wParam windows.WPARAM, lParam windows.LPARAM) windows.LRESULT {
			if code == user32.HC_ACTION {
				p := (*user32.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
				kb.Events <- KeyboardEvent{Key: uint32(p.VkCode), State: wParam == user32.WM_KEYDOWN}
			}
			return user32.CallNextHookEx(0, code, wParam, lParam)
		},
		0,
		0,
	)
	defer user32.UnhookWindowsHookEx(hook)

	var msg user32.MSG
	for user32.GetMessageW(&msg, 0, 0, 0) == 1 && !kb.Closed() {
		user32.TranslateMessage(&msg)
		user32.DispatchMessageW(&msg)
	}
}

func getState(key int) bool {
	return user32.GetAsyncKeyState(key) != 0
}

func sendInput(key int, state bool) {
	if state {
		user32.Keybd_event(windows.BYTE(key), 0, 0, 0)
		return
	}
	user32.Keybd_event(windows.BYTE(key), 0, user32.KEYEVENTF_KEYUP, 0)
}
