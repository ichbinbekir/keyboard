package keyboard

func MouseEvent(event uintptr) error {
	if _, _, err := _mouse_event.Call(event); err.Error() != errOperationComletedText {
		return err
	}

	return nil
}
