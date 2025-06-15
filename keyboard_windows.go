package keyboard

import (
	"unsafe"

	"github.com/ichbinbekir/keyboard/pkg/windows/core"
	"github.com/ichbinbekir/keyboard/pkg/windows/user32"
)

func (kb *Keyboard) readEvents() {
	hook := user32.SetWindowsHookExW(
		user32.WH_KEYBOARD_LL,
		func(code int, wParam core.WPARAM, lParam core.LPARAM) core.LRESULT {
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
	for user32.GetMessageW(&msg, 0, 0, 0) == 1 && !kb.IsClosed() {
		user32.TranslateMessage(&msg)
		user32.DispatchMessageW(&msg)
	}
}

func getState(key int) bool {
	return user32.GetAsyncKeyState(key) != 0
}
