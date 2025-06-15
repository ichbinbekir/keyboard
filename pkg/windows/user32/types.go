package user32

import "github.com/ichbinbekir/keyboard/pkg/windows/core"

type (
	POINT struct {
		X core.LONG
		Y core.LONG
	}
	MSG struct {
		Hwnd    core.HWND
		Message core.UINT
		WParam  core.WPARAM
		LParam  core.LPARAM
		Time    core.DWORD
		Pt      POINT
	}
	KBDLLHOOKSTRUCT struct {
		VkCode      core.DWORD
		ScanCode    core.DWORD
		Flags       core.DWORD
		Time        core.DWORD
		DwExtraInfo core.ULONG_PTR
	}
	MSLLHOOKSTRUCT struct {
		Pt          POINT
		MouseData   core.DWORD
		Flags       core.DWORD
		Time        core.DWORD
		DwExtraInfo core.ULONG_PTR
	}
)

type HOOKPROC func(code int, wParam core.WPARAM, lParam core.LPARAM) core.LRESULT
