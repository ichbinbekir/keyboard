package user32

import (
	"syscall"
	"unsafe"

	"github.com/ichbinbekir/keyboard/pkg/windows/core"
)

func SetWindowsHookExW(idHook int, lpfn HOOKPROC, hmod core.HINSTANCE, dwThreadId core.DWORD) core.HHOOK {
	ret, _, _ := _setWindowsHookExW.Call(
		uintptr(idHook),
		syscall.NewCallback(lpfn),
		uintptr(hmod),
		uintptr(dwThreadId),
	)
	return core.HHOOK(ret)
}

func GetMessageW(lpMsg *MSG, hWnd core.HWND, wMsgFilterMin, wMsgFilterMax core.UINT) core.BOOL {
	ret, _, _ := _getMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return core.BOOL(ret)
}

func TranslateMessage(lpMsg *MSG) core.BOOL {
	ret, _, _ := _translateMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return core.BOOL(ret)
}

func DispatchMessageW(lpMsg *MSG) core.LRESULT {
	ret, _, _ := _dispatchMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return core.LRESULT(ret)
}

func UnhookWindowsHookEx(hhk core.HHOOK) core.BOOL {
	ret, _, _ := _unhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return core.BOOL(ret)
}

func CallNextHookEx(hhk core.HHOOK, nCode int, wParam core.WPARAM, lParam core.LPARAM) core.LRESULT {
	ret, _, _ := _callNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return core.LRESULT(ret)
}

func GetAsyncKeyState(vKey int) core.SHORT {
	ret, _, _ := _getAsyncKeyState.Call(
		uintptr(vKey),
	)
	return core.SHORT(ret)
}

func Keybd_event(bVk, bScan core.BYTE, dwFlags core.DWORD, dwExtraInfo core.ULONG_PTR) {
	_keybd_event.Call(
		uintptr(bVk),
		uintptr(bScan),
		uintptr(dwFlags),
		uintptr(dwExtraInfo),
	)
}
