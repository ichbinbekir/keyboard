package keyboard

import (
	"errors"
	"unsafe"

	"github.com/ichbinbekir/windows"
	"github.com/ichbinbekir/windows/user32"
)

func (kb *Keyboard) readEvents() {
	handleMouse := kb.config.HandleMouseButtons || kb.config.HandleMouseWheel || kb.config.HandleMouseMove
	if !kb.config.HandleKeyboard && !handleMouse {
		kb.Errors <- errors.New("you must handle keyboard or mouse")
		return
	}

	if kb.config.HandleKeyboard {
		hook := user32.SetWindowsHookExW(user32.WH_KEYBOARD_LL, func(code int, wParam windows.WPARAM, lParam windows.LPARAM) windows.LRESULT {
			if code == user32.HC_ACTION {
				hookData := (*user32.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
				vkCode := uint32(hookData.VkCode)

				isDown := wParam == user32.WM_KEYDOWN /*|| wParam == user32.WM_SYSKEYDOWN*/

				if !kb.config.SendRepeatedKeyDowns && isDown && kb.keyStates[vkCode] {
					// It's a repeat, and we're not sending repeats, so just return.
					return user32.CallNextHookEx(0, code, wParam, lParam)
				}

				kb.keyStates[vkCode] = isDown
				kb.Events <- ButtonEvent{KeyCode: vkCode, KeyState: isDown}
			}
			return user32.CallNextHookEx(0, code, wParam, lParam)
		}, 0, 0)
		defer user32.UnhookWindowsHookEx(hook)
	}

	if handleMouse {
		hook := user32.SetWindowsHookExW(user32.WH_MOUSE_LL, func(code int, wParam windows.WPARAM, lParam windows.LPARAM) windows.LRESULT {
			if code == user32.HC_ACTION {
				hookData := (*user32.MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))

				if kb.config.HandleMouseButtons {
					var keyCode uint32
					var state bool
					switch wParam {
					case user32.WM_LBUTTONDOWN:
						keyCode = user32.VK_LBUTTON
						state = true
					case user32.WM_LBUTTONUP:
						keyCode = user32.VK_LBUTTON
					case user32.WM_RBUTTONDOWN:
						keyCode = user32.VK_RBUTTON
						state = true
					case user32.WM_RBUTTONUP:
						keyCode = user32.VK_RBUTTON
					case user32.WM_MBUTTONDOWN:
						keyCode = user32.VK_MBUTTON
						state = true
					case user32.WM_MBUTTONUP:
						keyCode = user32.VK_MBUTTON
					case user32.WM_XBUTTONDOWN:
						// The high-order word of MouseData contains the XButton that was pressed.
						// 1 = XBUTTON1, 2 = XBUTTON2
						xbutton := hookData.MouseData >> 16
						switch xbutton {
						case 1:
							keyCode = user32.VK_XBUTTON1
						case 2:
							keyCode = user32.VK_XBUTTON2
						}
						state = true
					case user32.WM_XBUTTONUP:
						// The high-order word of MouseData contains the XButton that was released.
						// 1 = XBUTTON1, 2 = XBUTTON2
						xbutton := hookData.MouseData >> 16
						switch xbutton {
						case 1:
							keyCode = user32.VK_XBUTTON1
						case 2:
							keyCode = user32.VK_XBUTTON2
						}
					}
					if keyCode != 0 {
						kb.Events <- ButtonEvent{KeyCode: keyCode, KeyState: state}
					}
				}

				if kb.config.HandleMouseWheel {
					if wParam == user32.WM_MOUSEWHEEL {
						// The high-order word of MouseData is the wheel delta.
						// A positive value indicates that the wheel was rotated forward.
						// A negative value indicates that the wheel was rotated backward.
						delta := int16(hookData.MouseData >> 16)
						kb.Events <- WheelEvent{Delta: delta}
					}
				}

				if kb.config.HandleMouseMove {
					if hookData.Pt.X != kb.lastMousePos.X || hookData.Pt.Y != kb.lastMousePos.Y {
						kb.lastMousePos.X = hookData.Pt.X
						kb.lastMousePos.Y = hookData.Pt.Y
						kb.Events <- MouseMoveEvent{X: int32(hookData.Pt.X), Y: int32(hookData.Pt.Y)}
					}
				}
			}
			return user32.CallNextHookEx(0, code, wParam, lParam)
		}, 0, 0)
		defer user32.UnhookWindowsHookEx(hook)
	}

	var msg user32.MSG
	for user32.GetMessageW(&msg, 0, 0, 0) == 1 {
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
